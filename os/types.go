package os

import (
	ofs "io/fs"
	oos "os"

	"github.com/avfs/avfs"
)

type DirEntry oos.DirEntry
type FileInfo oos.FileInfo
type FileMode ofs.FileMode
type LinkError oos.LinkError
type PathError oos.PathError
type ProcAttr oos.ProcAttr
type Process oos.Process
type ProcessState oos.ProcessState
type Signal oos.Signal
type SyscallError oos.SyscallError
type File avfs.File
