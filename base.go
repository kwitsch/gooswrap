package gooswrap

import (
	"github.com/avfs/avfs"
	"github.com/avfs/avfs/vfs/memfs"
	"github.com/avfs/avfs/vfs/osfs"
)

// data store for wrapper object
type WrapperStore struct {
	// filesystem wrapper
	Fs avfs.VFS
	// virtual data store
	Virtual *VirtualData
}

// wrapper object
var Wrapper *WrapperStore

// nolint:gochecknoinits
func init() {
	newWrapper(false)
}

// initialize new virtual wrapper
func NewVirtual() {
	newWrapper(true)
}

// set wrapper to os mode
func ToOs() {
	if Wrapper.IsVirtual() {
		newWrapper(false)
	}
}

// sets current wrapper to new one
func newWrapper(virtual bool) {
	var fs avfs.VFS
	if virtual {
		fs = memfs.New(memfs.WithMainDirs())
	} else {
		fs = osfs.New()
	}

	wrapper := WrapperStore{
		Fs: fs,
	}

	if virtual {
		wrapper.Virtual = newVirtualData()
	}

	Wrapper = &wrapper
}

// Is the Wrapper virtual
func (ws *WrapperStore) IsVirtual() bool {
	return (ws.Virtual != nil)
}
