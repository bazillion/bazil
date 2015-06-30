package server

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"bazil.org/bazil/cas/chunks/kvchunks"
	"bazil.org/bazil/db"
	"bazil.org/bazil/fs"
	"bazil.org/bazil/kv"
	"bazil.org/bazil/kv/kvfiles"
	"bazil.org/bazil/kv/kvmulti"
	"bazil.org/bazil/kv/kvpeer"
	"bazil.org/bazil/kv/untrusted"
	"bazil.org/bazil/peer"
	"bazil.org/bazil/tokens"
	"bazil.org/fuse"
	fusefs "bazil.org/fuse/fs"
	"github.com/boltdb/bolt"
)

type VolumeRef struct {
	// closed after the serve loop has exited
	unmounted chan struct{}
}

type App struct {
	DataDir  string
	lockFile *os.File
	DB       *db.DB
	volumes  struct {
		sync.Mutex
		open map[db.VolumeID]*VolumeRef
	}
	Keys *CryptoKeys
	tls  struct {
		config atomic.Value
		gen    sync.Mutex
	}
}

func New(dataDir string) (app *App, err error) {
	err = os.Mkdir(dataDir, 0700)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	lockPath := filepath.Join(dataDir, "lock")
	lockFile, err := lock(lockPath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			// if we're reporting an error, also unlock
			_ = lockFile.Close()
		}
	}()

	kvpath := filepath.Join(dataDir, "chunks")
	err = kvfiles.Create(kvpath)
	if err != nil {
		return nil, err
	}

	dbpath := filepath.Join(dataDir, "bazil.bolt")
	database, err := db.Open(dbpath, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = database.DB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(tokens.BucketBazil)); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		database.Close()
		return nil, err
	}

	keys, err := loadOrGenerateKeys(database.DB)
	if err != nil {
		return nil, err
	}

	app = &App{
		DataDir:  dataDir,
		lockFile: lockFile,
		DB:       database,
		Keys:     keys,
	}
	app.volumes.open = make(map[db.VolumeID]*VolumeRef)
	return app, nil
}

func (app *App) Close() {
	app.DB.Close()
	app.lockFile.Close()
}

// TODO this function smells
func (app *App) serveMount(vol *fs.Volume, id *db.VolumeID, mountpoint string, conn *fuse.Conn, srv *fusefs.Server) error {
	serveErr := make(chan error, 1)
	go func() {
		defer func() {
			// remove map entry on unmount or failed mount
			app.volumes.Lock()
			delete(app.volumes.open, *id)
			app.volumes.Unlock()
		}()
		defer conn.Close()
		serveErr <- srv.Serve(vol)
	}()

	select {
	case <-conn.Ready:
		if err := conn.MountError; err != nil {
			return fmt.Errorf("mount fail (delayed): %v", err)
		}
		return nil
	case err := <-serveErr:
		// Serve quit early
		if err != nil {
			return fmt.Errorf("filesystem failure: %v", err)
		}
		return errors.New("Serve exited early")
	}
}

type MountInfo struct {
	VolumeID db.VolumeID
}

func (app *App) openKV(tx *db.Tx, storage *db.VolumeStorage) (kv.KV, error) {
	var kvstores []kv.KV

	c := storage.Cursor()
	for item := c.First(); item != nil; item = c.Next() {
		backend, err := item.Backend()
		if err != nil {
			return nil, err
		}
		s, err := app.openStorage(backend)
		if err != nil {
			return nil, err
		}

		sharingKeyName, err := item.SharingKeyName()
		if err != nil {
			return nil, err
		}
		sharingKey, err := tx.SharingKeys().Get(sharingKeyName)
		if err != nil {
			return nil, fmt.Errorf("getting sharing key %q: %v", sharingKeyName, err)
		}
		var secret [32]byte
		sharingKey.Secret(&secret)
		s = untrusted.New(s, &secret)

		kvstores = append(kvstores, s)
	}

	return kvmulti.New(kvstores...), nil
}

func (app *App) openStorage(backend string) (kv.KV, error) {
	switch backend {
	case "local":
		kvpath := filepath.Join(app.DataDir, "chunks")
		return kvfiles.Open(kvpath)
	}
	if backend != "" && backend[0] == '/' {
		return kvfiles.Open(backend)
	}
	idx := strings.IndexByte(backend, ':')
	if idx != -1 {
		scheme, rest := backend[:idx], backend[idx+1:]
		switch scheme {
		case "peerkey":
			var key peer.PublicKey
			if err := key.Set(rest); err != nil {
				return nil, err
			}
			p, err := app.DialPeer(&key)
			if err != nil {
				return nil, err
			}
			// TODO Close
			return kvpeer.Open(p)
		}
	}
	return nil, errors.New("unknown storage backend")
}

func (app *App) ValidateKV(backend string) error {
	// TODO opening a KV just to validate the input string would be a
	// bad idea if a backend was costly to open; then again, current
	// API doesn't support a Close anyway.
	_, err := app.openStorage(backend)
	return err
}

// Mount makes the contents of the volume visible at the given
// mountpoint. If Mount returns with a nil error, the mount has
// occurred.
func (app *App) Mount(volumeName string, mountpoint string) (*MountInfo, error) {
	// TODO obey `bazil -debug server run`

	info := &MountInfo{}
	var ready = make(chan error, 1)
	app.volumes.Lock()
	err := app.DB.View(func(tx *db.Tx) error {
		v, err := tx.Volumes().GetByName(volumeName)
		if err != nil {
			return err
		}
		v.VolumeID(&info.VolumeID)

		if _, ok := app.volumes.open[info.VolumeID]; ok {
			return errors.New("volume already mounted")
		}

		conn, err := fuse.Mount(mountpoint,
			fuse.MaxReadahead(32*1024*1024),
			fuse.AsyncRead(),
		)
		if err != nil {
			return fmt.Errorf("mount fail: %v", err)
		}

		srv := fusefs.New(conn, nil)

		kvstore, err := app.openKV(tx, v.Storage())
		if err != nil {
			return err
		}

		chunkStore := kvchunks.New(kvstore)
		vol, err := fs.Open(app.DB, chunkStore, &info.VolumeID, (*peer.PublicKey)(app.Keys.Sign.Pub))
		if err != nil {
			return err
		}
		mnt := &VolumeRef{
			unmounted: make(chan struct{}),
		}
		go func() {
			defer close(mnt.unmounted)
			ready <- app.serveMount(vol, &info.VolumeID, mountpoint, conn, srv)
		}()
		app.volumes.open[info.VolumeID] = mnt
		return nil
	})
	app.volumes.Unlock()
	if err != nil {
		return nil, err
	}
	err = <-ready
	if err != nil {
		return nil, err
	}
	return info, nil
}

var ErrNotMounted = errors.New("not currently mounted")

func (app *App) WaitForUnmount(volumeID *db.VolumeID) error {
	app.volumes.Lock()
	// we hold onto mnt after releasing the lock, but it's safe in
	// this case; gc keeps it pinned, and we don't look at mutable
	// data
	mnt, ok := app.volumes.open[*volumeID]
	app.volumes.Unlock()
	if !ok {
		return ErrNotMounted
	}
	<-mnt.unmounted
	return nil
}
