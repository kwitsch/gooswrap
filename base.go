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
	// virtual data store
	Virtual *Virtual
}

// working mode indicator and data store
type Virtual struct {
	// is the current wrapper in virtual mode
	virtual *bool
	// current working directory
	WorkingDirectory string
	// data store for virtual mode
	Data *VirtualData
}

// data store for virtual mode
type VirtualData struct {
	// pseudo environment
	Env map[string]string
	// editable hostname
	Hostname string
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

	util := afero.Afero{
		Fs: fs,
	}

	vs := Virtual{
		virtual: &virtual,
	}

	if virtual {
		vs.Data = newVirtualData()
	}

	Wrapper = &WrapperStore{
		Fs:      fs,
		Util:    &util,
		Virtual: &vs,
	}
}

// returns new virtual data store
func newVirtualData() *VirtualData {
	hostname := "virtual"
	if thn, err := oos.Hostname(); err == nil {
		hostname = thn
	}
	return &VirtualData{
		Env:      make(map[string]string),
		Hostname: hostname,
	}
}

// Is the Wrapper virtual
func (ws *WrapperStore) IsVirtual() bool {
	return ws.Virtual.IsVirtual()
}

// is the Virtual object virtual
func (v *Virtual) IsVirtual() bool {
	return *v.virtual
}

// get file path prefixed with current working directory if it doesen't start with a slash
func (ws *WrapperStore) GetPath(fpath string) string {
	if strings.HasPrefix(fpath, "/") {
		return fpath
	} else {
		return path.Join(ws.Virtual.WorkingDirectory, fpath)
	}
}

// sync current os environment variables to virtual environment variables
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *Virtual) SyncEnv() error {
	return v.onlyWhenVirtual(func() {
		oenv := oos.Environ()
		for _, ec := range oenv {
			es := strings.Split(ec, "=")
			if len(es) == 2 {
				v.Data.Env[es[0]] = es[1]
			}
		}
	})
}

// set virtual hostname
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *Virtual) SetHostname(hostname string) error {
	return v.onlyWhenVirtual(func() {
		Wrapper.Virtual.Data.Hostname = hostname
	})
}

// create TempDir, UserCacheDir, UserConfigDir & UserHomeDir
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *Virtual) InitDirectories() error {
	return v.onlyWhenVirtualReturn(func() error {
		if err := Wrapper.Fs.MkdirAll(oos.TempDir(), 0777); err != nil {
			return err
		}
		if ucd, err := oos.UserCacheDir(); err != nil {
			if err := Wrapper.Fs.MkdirAll(ucd, 0777); err != nil {
				return err
			}
		}
		if ucd, err := oos.UserConfigDir(); err != nil {
			if err := Wrapper.Fs.MkdirAll(ucd, 0777); err != nil {
				return err
			}
		}
		if uhd, err := oos.UserHomeDir(); err != nil {
			if err := Wrapper.Fs.MkdirAll(uhd, 0777); err != nil {
				return err
			}
		}
		return nil
	})
}

// executes action if in virtual mode
// returns ErrNotVirtual if not in virtual mode
func (v *Virtual) onlyWhenVirtual(action func()) error {
	if v.IsVirtual() {
		action()
		return nil
	}

	return ErrNotVirtual
}

// executes action if in virtual mode and returns its error
// returns ErrNotVirtual if not in virtual mode
func (v *Virtual) onlyWhenVirtualReturn(action func() error) error {
	if v.IsVirtual() {
		return action()
	}

	return ErrNotVirtual
}
