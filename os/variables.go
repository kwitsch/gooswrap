package os

import (
	"errors"
	oos "os"

	"github.com/spf13/afero"
)

var (
	// ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	ErrInvalid = oos.ErrInvalid // "invalid argument"

	ErrPermission = oos.ErrPermission // "permission denied"
	ErrExist      = oos.ErrExist      // "file already exists"
	ErrNotExist   = oos.ErrNotExist   // "file does not exist"
	ErrClosed     = oos.ErrClosed     // "file already closed"

	ErrNoDeadline       = oos.ErrNoDeadline       // "file type does not support deadline"
	ErrDeadlineExceeded = oos.ErrDeadlineExceeded // "i/o timeout"
)

var (
	Stdin  = oos.Stdin
	Stdout = oos.Stdout
	Stderr = oos.Stderr
)

var Args []string

var ErrProcessDone = oos.ErrProcessDone

var (
	ErrFileClosed = afero.ErrFileClosed
	ErrOutOfRange = afero.ErrOutOfRange
	ErrTooLarge   = afero.ErrTooLarge
	ErrNoReadlink = afero.ErrNoReadlink
	ErrNoSymlink  = afero.ErrNoSymlink
	ErrNoLink     = errors.New("link not supported")
	ErrNoPipe     = errors.New("pipe not supported")
)
