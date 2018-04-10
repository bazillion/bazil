// Package db contains a database abstraction layer.
package db // import "bazil.org/bazil/db"

import (
	"os"

	"github.com/boltdb/bolt"
)

// DB provides abstracted access to the Bolt database used by the
// server.
type DB struct {
	*bolt.DB
}

func Open(path string, mode os.FileMode, options *bolt.Options) (*DB, error) {
	d, err := bolt.Open(path, mode, options)
	if err != nil {
		return nil, err
	}
	db := &DB{d}
	if err := db.Update(db.init); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// init sets up the initial database contents. It is guaranteed to be
// idempotent and safe to run on pre-existing databases.
func (db *DB) init(tx *Tx) error {
	if err := tx.initVolumes(); err != nil {
		return err
	}
	if err := tx.initPeers(); err != nil {
		return err
	}
	if err := tx.initSharingKeys(); err != nil {
		return err
	}
	return nil
}

func (db *DB) View(fn func(*Tx) error) error {
	wrapper := func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	}
	return db.DB.View(wrapper)
}

// Update makes changes to the database. There can be only one Update
// call at a time.
//
// If a lock L is held while calling db.Update, L must never be taken
// inside a write transaction, at the risk of a deadlock.
func (db *DB) Update(fn func(*Tx) error) error {
	wrapper := func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	}
	return db.DB.Update(wrapper)
}

// Tx is a database transaction.
//
// Unless otherwise stated, any values returned by methods here (and
// transitively from their methods) are only valid while the
// transaction is alive. This does not apply to returned error values.
type Tx struct {
	*bolt.Tx
}
