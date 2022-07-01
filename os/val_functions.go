package os

import (
	oos "os"

	. "github.com/kwitsch/gooswrap"
)

func Getegid() int {
	return getIntVal(oos.Getegid, Wrapper.Virtual.Egid)
}

func Geteuid() int {
	return getIntVal(oos.Geteuid, Wrapper.Virtual.Euid)
}

func Getgid() int {
	return getIntVal(oos.Getgid, Wrapper.Virtual.Gid)
}

func Getgroups() ([]int, error) {
	if Wrapper.IsVirtual() {
		return Wrapper.Virtual.Groups()
	}
	return oos.Getgroups()
}

func Getpagesize() int {
	return getIntVal(oos.Getpagesize, Wrapper.Virtual.Pagesize)
}

func Getpid() int {
	return getIntVal(oos.Getpid, Wrapper.Virtual.Pid)
}

func Getppid() int {
	return getIntVal(oos.Getppid, Wrapper.Virtual.Ppid)
}

func Getuid() int {
	return getIntVal(oos.Getuid, Wrapper.Virtual.Uid)
}

func Getwd() (string, error) {
	return Wrapper.Fs.Getwd()
}

func Hostname() (string, error) {
	if !Wrapper.IsVirtual() {
		return oos.Hostname()
	} else {
		return Wrapper.Virtual.Hostname()
	}
}

func getIntVal(of func() int, vf func() (int, error)) int {
	if Wrapper.IsVirtual() {
		if res, err := vf(); err == nil {
			return res
		}
	}

	return of()
}
