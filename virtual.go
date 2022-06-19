package gooswrap

import (
	oos "os"
	"strings"
)

// data store for virtual mode
type VirtualData struct {
	// pseudo environment
	Env map[string]string
	// editable hostname
	Hostname string
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

// sync current os environment variables to virtual environment variables
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SyncEnv() error {
	return onlyWhenVirtual(v, func() {
		oenv := oos.Environ()
		for _, ec := range oenv {
			es := strings.Split(ec, "=")
			if len(es) == 2 {
				v.Env[es[0]] = es[1]
			}
		}
	})
}

// set virtual hostname
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetHostname(hostname string) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.Hostname = hostname
	})
}

// create TempDir, UserCacheDir, UserConfigDir & UserHomeDir
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) InitDirectories() error {
	return onlyWhenVirtualError(v, func() error {
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
func onlyWhenVirtual(v *VirtualData, action func()) error {
	if v != nil {
		action()
		return nil
	}

	return ErrNotVirtual
}

// executes action if in virtual mode and returns its error
// returns ErrNotVirtual if not in virtual mode
func onlyWhenVirtualError(v *VirtualData, action func() error) error {
	if v != nil {
		return action()
	}

	return ErrNotVirtual
}
