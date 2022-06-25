package gooswrap

import (
	oos "os"
	"strings"
)

const defaultHostname string = "virtual"

// data store for virtual mode
type VirtualData struct {
	egid     int
	euid     int
	gid      int
	groups   []int
	pagesize int
	pid      int
	ppid     int
	uid      int
	// editable hostname
	hostname string
	// pseudo environment
	Env map[string]string
}

// returns new virtual data store
func newVirtualData() *VirtualData {
	return &VirtualData{
		egid:     oos.Getegid(),
		euid:     oos.Geteuid(),
		gid:      oos.Getgid(),
		groups:   make([]int, 0),
		pagesize: oos.Getpagesize(),
		pid:      oos.Getpid(),
		ppid:     oos.Getppid(),
		uid:      oos.Getuid(),
		hostname: defaultHostname,
		Env:      make(map[string]string),
	}
}

// set virtual Egid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetEgid(egid int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.egid = egid
	})
}

// get virtual Egid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Egid() (int, error) {
	return onlyWhenVirtualIntError(v, func() (int, error) {
		return v.egid, nil
	})
}

// set virtual Euid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetEuid(euid int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.euid = euid
	})
}

// get virtual Euid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Euid() (int, error) {
	return onlyWhenVirtualIntError(v, func() (int, error) {
		return v.euid, nil
	})
}

// set virtual Gid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetGid(gid int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.gid = gid
	})
}

// get virtual Gid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Gid() (int, error) {
	return onlyWhenVirtualIntError(v, func() (int, error) {
		return v.gid, nil
	})
}

// set virtual Groups
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetGroups(groups []int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.groups = groups
	})
}

// get virtual Groups
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Groups() ([]int, error) {
	if v != nil {
		return v.groups, nil
	}
	return make([]int, 0), ErrNotVirtual
}

// set virtual Pagesize
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetPagesize(pagesize int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.pagesize = pagesize
	})
}

// get virtual Pagesize
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Pagesize() (int, error) {
	return onlyWhenVirtualIntError(v, func() (int, error) {
		return v.pagesize, nil
	})
}

// set virtual Pid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetPid(pid int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.pid = pid
	})
}

// get virtual Pid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Pid() (int, error) {
	return onlyWhenVirtualIntError(v, func() (int, error) {
		return v.pid, nil
	})
}

// set virtual Ppid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetPpid(ppid int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.ppid = ppid
	})
}

// get virtual Ppid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Ppid() (int, error) {
	return onlyWhenVirtualIntError(v, func() (int, error) {
		return v.ppid, nil
	})
}

// set virtual Uid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetUid(uid int) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.uid = uid
	})
}

// get virtual Uid
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Uid() (int, error) {
	return onlyWhenVirtualIntError(v, func() (int, error) {
		return v.uid, nil
	})
}

// set virtual hostname
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) SetHostname(hostname string) error {
	return onlyWhenVirtual(v, func() {
		Wrapper.Virtual.hostname = hostname
	})
}

// get virtual hostname
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) Hostname() (string, error) {
	return onlyWhenVirtualStringError(v, func() (string, error) {
		return v.hostname, nil
	})
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

// create TempDir, UserCacheDir, UserConfigDir & UserHomeDir
// returns ErrNotVirtual if Wrapper isen't in virtual mode
func (v *VirtualData) InitDirectories() error {
	return onlyWhenVirtualError(v, func() error {
		if err := Wrapper.Fs.MkdirAll(VirtualTempDir, 0777); err != nil {
			return err
		}
		if err := Wrapper.Fs.MkdirAll(VirtualUserCacheDir, 0777); err != nil {
			return err
		}
		if err := Wrapper.Fs.MkdirAll(VirtualUserConfigDir, 0777); err != nil {
			return err
		}
		if err := Wrapper.Fs.MkdirAll(VirtualUserHomeDir, 0777); err != nil {
			return err
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

// executes action if in virtual mode and returns its error
// returns ErrNotVirtual if not in virtual mode
func onlyWhenVirtualStringError(v *VirtualData, action func() (string, error)) (string, error) {
	if v != nil {
		return action()
	}

	return "", ErrNotVirtual
}

// executes action if in virtual mode and returns its error
// returns ErrNotVirtual if not in virtual mode
func onlyWhenVirtualIntError(v *VirtualData, action func() (int, error)) (int, error) {
	if v != nil {
		return action()
	}

	return 0, ErrNotVirtual
}
