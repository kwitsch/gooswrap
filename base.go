package gooswrap

import (
	"github.com/avfs/avfs"
	"github.com/avfs/avfs/idm/memidm"
	"github.com/avfs/avfs/idm/osidm"
	"github.com/avfs/avfs/vfs/memfs"
	"github.com/avfs/avfs/vfs/osfs"
)

// data store for wrapper object
type WrapperStore struct {
	// filesystem wrapper
	Fs avfs.VFS
	// Identity manager
	Idm avfs.IdentityMgr
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
	Wrapper.clear()

	var idm avfs.IdentityMgr
	if virtual {
		idm = memidm.New()
	} else {
		idm = osidm.New()
	}

	var fs avfs.VFS
	if virtual {
		fs = memfs.New(memfs.WithMainDirs(), memfs.WithIdm(idm))
	} else {
		fs = osfs.New()
	}

	wrapper := WrapperStore{
		Fs:  fs,
		Idm: idm,
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

// clears virtual directory
func (ws *WrapperStore) clear() {
	if ws.IsVirtual() {
		ws.Fs.RemoveAll("/")
	}
}
