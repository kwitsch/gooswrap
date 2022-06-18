package gooswrap

import (
	oos "os"
	"path"
	"strings"

	"github.com/spf13/afero"
)

type WrapperStore struct {
	Fs      afero.Fs
	Util    *afero.Afero
	Virtual *Virtual
}

type Virtual struct {
	virtual *bool
	Data    *VirtualData
}

type VirtualData struct {
	Env      map[string]string
	Hostname string
	Path     string
}

var Wrapper *WrapperStore

func init() {
	newWrapper(false)
}

func ToVirtual() {
	newWrapper(true)
}

func ToOs() {
	if Wrapper.IsVirtual() {
		newWrapper(false)
	}
}

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

	v := Virtual{
		virtual: &virtual,
	}

	if virtual {
		v.Data = newVirtualData()
	}

	Wrapper = &WrapperStore{
		Fs:      fs,
		Util:    &util,
		Virtual: &v,
	}
}

func newVirtualData() *VirtualData {
	hostname := "virtual"
	if thn, err := oos.Hostname(); err == nil {
		hostname = thn
	}

	return &VirtualData{
		Env:      make(map[string]string),
		Hostname: hostname,
		Path:     "/",
	}
}

func (ws *WrapperStore) IsVirtual() bool {
	return *Wrapper.Virtual.virtual
}

func (ws *WrapperStore) GetPath(fpath string) string {
	if strings.HasPrefix(fpath, "/") {
		return fpath
	} else {
		return path.Join(ws.Virtual.Data.Path, fpath)
	}
}

func (ws *WrapperStore) SyncEnv() error {
	return ws.onlyWhenVirtual(func() {
		oenv := oos.Environ()
		for _, ec := range oenv {
			es := strings.Split(ec, "=")
			if len(es) == 2 {
				Wrapper.Virtual.Data.Env[es[0]] = es[1]
			}
		}
	})
}

func (ws *WrapperStore) SetHostname(hostname string) error {
	return ws.onlyWhenVirtual(func() {
		Wrapper.Virtual.Data.Hostname = hostname
	})
}

func (ws *WrapperStore) onlyWhenVirtual(action func()) error {
	if ws.IsVirtual() {
		action()
		return nil
	}

	return ErrNotVirtual
}
