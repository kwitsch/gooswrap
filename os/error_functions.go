package os

import (
	oos "os"
)

// IsExist returns a boolean indicating whether the error is known to report
// that a file or directory already exists. It is satisfied by ErrExist as
// well as some syscall errors.
//
// This function predates errors.Is. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrExist).
func IsExist(err error) bool {
	return oos.IsExist(err)
}

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
//
// This function predates errors.Is. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrNotExist).
func IsNotExist(err error) bool {
	return oos.IsNotExist(err)
}

// IsPermission returns a boolean indicating whether the error is known to
// report that permission is denied. It is satisfied by ErrPermission as well
// as some syscall errors.
//
// This function predates errors.Is. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrPermission).
func IsPermission(err error) bool {
	return oos.IsPermission(err)
}

// IsTimeout returns a boolean indicating whether the error is known
// to report that a timeout occurred.
//
// This function predates errors.Is, and the notion of whether an
// error indicates a timeout can be ambiguous. For example, the Unix
// error EWOULDBLOCK sometimes indicates a timeout and sometimes does not.
// New code should use errors.Is with a value appropriate to the call
// returning the error, such as os.ErrDeadlineExceeded.
func IsTimeout(err error) bool {
	return oos.IsTimeout(err)
}
