package os

import (
	"errors"
	"io/fs"
	oos "os"
	"syscall"
)

var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = fs.ErrInvalid // "invalid argument"

	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"

	ErrNoDeadline       = oos.ErrNoDeadline       // "file type does not support deadline"
	ErrDeadlineExceeded = oos.ErrDeadlineExceeded // "i/o timeout"
)

var (
	Stdin  = oos.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = oos.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = oos.NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

var Args []string

var ErrProcessDone = errors.New("os: process already finished")
