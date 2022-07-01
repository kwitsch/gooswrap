package gooswrap

import (
	oos "os"
	"path"
	"regexp"
	"strings"

	"github.com/avfs/avfs"
	"github.com/avfs/avfs/vfs/memfs"
	"github.com/avfs/avfs/vfs/osfs"
)

// data store for wrapper object
type WrapperStore struct {
	// filesystem wrapper
	Fs avfs.VFS
	// current working directory
	WorkingDirectory string
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

	wd := "/"
	if !virtual {
		if dir, err := oos.Getwd(); err == nil {
			wd = dir
		}
	}

	wrapper := WrapperStore{
		Fs:               fs,
		WorkingDirectory: wd,
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

// get file path prefixed with current working directory if it doesen't start with a slash
func (ws *WrapperStore) GetPath(fpath string) string {
	match, _ := regexp.MatchString("^[A-Z]:.*", fpath)
	if match || strings.HasPrefix(fpath, "/") {
		return fpath
	} else {
		return path.Join(ws.WorkingDirectory, fpath)
	}
}
