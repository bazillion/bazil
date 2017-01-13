package wire

import (
	"bazil.org/fuse"
)

// GetFUSEDirent returns a populated fuse.Dirent
func (d *Dirent) GetFUSEDirent(name string) fuse.Dirent {
	fde := fuse.Dirent{
		Inode: d.Inode,
		Name:  name,
	}
	switch {
	case d.GetFile() != nil:
		fde.Type = fuse.DT_File

	case d.GetDir() != nil:
		fde.Type = fuse.DT_Dir
	}
	return fde
}
