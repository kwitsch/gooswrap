package os

import (
	"io/fs"
	oos "os"
	"time"

	. "github.com/kwitsch/gooswrap"
)

func Chdir(dir string) error {
	if !Wrapper.IsVirtual() {
		err := oos.Chdir(dir)
		if err == nil {
			Wrapper.WorkingDirectory = dir
		}
		return err
	} else {
		exists, err := Wrapper.Util.DirExists(dir)
		if err != nil {
			return nil
		} else if !exists {
			return ErrNotExist
		} else {
			Wrapper.WorkingDirectory = dir
			return nil
		}
	}
}

func Chmod(name string, mode FileMode) error {
	return Wrapper.Fs.Chmod(Wrapper.GetPath(name), (fs.FileMode)(mode))
}

func Chown(name string, uid, gid int) error {
	return Wrapper.Fs.Chown(Wrapper.GetPath(name), uid, gid)
}

func Chtimes(name string, atime time.Time, mtime time.Time) error {
	return Wrapper.Fs.Chtimes(Wrapper.GetPath(name), atime, mtime)
}

func DirFS(dir string) fs.FS {
	return oos.DirFS(dir)
}

func Executable() (string, error) {
	return oos.Executable()
}

func Exit(code int) {
	oos.Exit(code)
}

func Expand(s string, mapping func(string) string) string {
	return oos.Expand(s, mapping)
}

func Lchown(name string, uid, gid int) error {
	if !Wrapper.IsVirtual() {
		return oos.Lchown(name, uid, gid)
	} else {
		return Wrapper.Fs.Chown(Wrapper.GetPath(name), uid, gid)
	}
}

func Link(oldname, newname string) error {
	if !Wrapper.IsVirtual() {
		return oos.Link(oldname, newname)
	} else {
		return ErrNotSupported
	}
}

func Mkdir(name string, perm FileMode) error {
	return Wrapper.Fs.Mkdir(Wrapper.GetPath(name), (oos.FileMode)(perm))
}

func MkdirAll(path string, perm FileMode) error {
	return Wrapper.Fs.MkdirAll(Wrapper.GetPath(path), (oos.FileMode)(perm))
}

func MkdirTemp(dir, pattern string) (string, error) {
	return oos.MkdirTemp(dir, pattern)
}

func NewSyscallError(syscall string, err error) error {
	return oos.NewSyscallError(syscall, err)
}

func Pipe() (*oos.File, *oos.File, error) {
	if !Wrapper.IsVirtual() {
		return oos.Pipe()
	} else {
		return nil, nil, ErrNotSupported
	}
}

func ReadFile(name string) ([]byte, error) {
	return Wrapper.Util.ReadFile(Wrapper.GetPath(name))
}

func Readlink(name string) (string, error) {
	if !Wrapper.IsVirtual() {
		return oos.Readlink(name)
	} else {
		return "", ErrNotSupported
	}
}

func Remove(name string) error {
	return Wrapper.Fs.Remove(Wrapper.GetPath(name))
}

func RemoveAll(path string) error {
	return Wrapper.Fs.RemoveAll(Wrapper.GetPath(path))
}

func Rename(oldpath, newpath string) error {
	return Wrapper.Fs.Rename(Wrapper.GetPath(oldpath), Wrapper.GetPath(newpath))
}

func SameFile(fi1, fi2 FileInfo) bool {
	return oos.SameFile(oos.FileInfo(fi1), oos.FileInfo(fi2))
}

func Symlink(oldname, newname string) error {
	if !Wrapper.IsVirtual() {
		return oos.Symlink(oldname, newname)
	} else {
		return ErrNotSupported
	}
}

func Truncate(name string, size int64) error {
	return oos.Truncate(name, size)
}

func WriteFile(name string, data []byte, perm FileMode) error {
	return Wrapper.Util.WriteFile(Wrapper.GetPath(name), data, (oos.FileMode)(perm))
}

func IsPathSeparator(c uint8) bool {
	return oos.IsPathSeparator(c)
}
