package gooswrap

import (
	oos "os"
	"path"
	"strings"

	"github.com/spf13/afero"
)

// data store for wrapper object
type WrapperStore struct {
	// afero filesystem
	Fs afero.Fs
	// afero ioutil
	Util *afero.Afero
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
	var fs afero.Fs
	if virtual {
		fs = afero.NewMemMapFs()
	} else {
		fs = afero.NewOsFs()
	}

	wd := "/"
	if !virtual {
		if dir, err := oos.Getwd(); err == nil {
			wd = dir
		}
	}

	util := afero.Afero{
		Fs: fs,
	}

	wrapper := WrapperStore{
		Fs:               fs,
		Util:             &util,
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
	if strings.HasPrefix(fpath, "/") {
		return fpath
	} else {
		return path.Join(ws.WorkingDirectory, fpath)
	}
}
