package os

import (
	oos "os"

	"github.com/spf13/afero"

	. "github.com/kwitsch/gooswrap"
)

type DirEntry oos.DirEntry
type FileInfo oos.FileInfo
type FileMode oos.FileMode
type LinkError oos.LinkError
type PathError oos.PathError
type ProcAttr oos.ProcAttr
type Process oos.Process
type ProcessState oos.ProcessState
type Signal oos.Signal
type SyscallError oos.SyscallError

type File struct {
	file afero.File
}

// Chdir changes the current working directory to the file,
// which must be a directory.
// If there is an error, it will be of type *PathError.
func (f *File) Chdir() error {
	return nil
}

// Chmod changes the mode of the file to mode.
// If there is an error, it will be of type *PathError.
func (f *File) Chmod(mode FileMode) error {
	return Wrapper.Fs.Chmod(f.file.Name(), (oos.FileMode)(mode))
}

// Chown changes the numeric uid and gid of the named file.
// If there is an error, it will be of type *PathError.
//
// On Windows, it always returns the syscall.EWINDOWS error, wrapped
// in *PathError.
func (f *File) Chown(uid, gid int) error {
	return Wrapper.Fs.Chown(f.file.Name(), uid, gid)
}

// Close closes the File, rendering it unusable for I/O.
// On files that support SetDeadline, any pending I/O operations will
// be canceled and return immediately with an ErrClosed error.
// Close will return an error if it has already been called.
func (f *File) Close() error {
	if f == nil {
		return ErrInvalid
	}
	return f.file.Close()
}
