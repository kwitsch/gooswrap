package os

import (
	"io/fs"
	oos "os"
	"time"

	. "github.com/kwitsch/gooswrap"
)

func Chdir(dir string) error {
	return Wrapper.Fs.Chdir(dir)
}

func Chmod(name string, mode FileMode) error {
	return Wrapper.Fs.Chmod(name, (fs.FileMode)(mode))
}

func Chown(name string, uid, gid int) error {
	return Wrapper.Fs.Chown(name, uid, gid)
}

func Chtimes(name string, atime time.Time, mtime time.Time) error {
	return Wrapper.Fs.Chtimes(name, atime, mtime)
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
	return Wrapper.Fs.Lchown(name, uid, gid)
}

func Link(oldname, newname string) error {
	return Wrapper.Fs.Link(oldname, newname)
}

func Mkdir(name string, perm FileMode) error {
	return Wrapper.Fs.Mkdir(name, (oos.FileMode)(perm))
}

func MkdirAll(path string, perm FileMode) error {
	return Wrapper.Fs.MkdirAll(path, (oos.FileMode)(perm))
}

func MkdirTemp(dir, pattern string) (string, error) {
	return Wrapper.Fs.MkdirTemp(dir, pattern)
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
	return Wrapper.Fs.ReadFile(name)
}

func Readlink(name string) (string, error) {
	return Wrapper.Fs.Readlink(name)
}

func Remove(name string) error {
	return Wrapper.Fs.Remove(name)
}

func RemoveAll(path string) error {
	return Wrapper.Fs.RemoveAll(path)
}

func Rename(oldpath, newpath string) error {
	return Wrapper.Fs.Rename(oldpath, newpath)
}

func SameFile(fi1, fi2 FileInfo) bool {
	return oos.SameFile(oos.FileInfo(fi1), oos.FileInfo(fi2))
}

func Symlink(oldname, newname string) error {
	return Wrapper.Fs.Symlink(oldname, newname)
}

func Truncate(name string, size int64) error {
	return Wrapper.Fs.Truncate(name, size)
}

func WriteFile(name string, data []byte, perm FileMode) error {
	return Wrapper.Fs.WriteFile(name, data, (oos.FileMode)(perm))
}

func IsPathSeparator(c uint8) bool {
	return oos.IsPathSeparator(c)
}
