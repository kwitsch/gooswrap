package gooswrap

import (
	"github.com/spf13/afero"
)

type WrapperStore struct {
	Fs      afero.Fs
	Util    *afero.Afero
	Virtual *bool
	Data    *VirtualData
}

type VirtualData struct {
	Env map[string]string
}

var Wrapper *WrapperStore

func init() {
	newWrapper(false)
}

func ToVirtual() {
	newWrapper(true)
}

func ToOs() {
	if *Wrapper.Virtual {
		newWrapper(false)
	}
}

func IsVirtual() bool {
	return *Wrapper.Virtual
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

	wrap := WrapperStore{
		Fs:      fs,
		Util:    &util,
		Virtual: &virtual,
	}

	if virtual {
		wrap.Data = newVirtualData()
	}

	Wrapper = &wrap
}

func newVirtualData() *VirtualData {
	res := VirtualData{
		Env: make(map[string]string),
	}

	return &res
}
