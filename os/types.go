package os

import (
	oos "os"

	"github.com/spf13/afero"
)

type DirEntry oos.DirEntry
type File afero.File
type FileInfo oos.FileInfo
type FileMode oos.FileMode
type LinkError oos.LinkError
type PathError oos.PathError
type ProcAttr oos.ProcAttr
type Process oos.Process
type ProcessState oos.ProcessState
type Signal oos.Signal
type SyscallError oos.SyscallError
