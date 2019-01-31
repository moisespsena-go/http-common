package httpcommon

import (
	"os"

	"github.com/moisespsena-go/os-common"
)

type Dir struct {
	Name   string
	Reader func(count int) (items []os.FileInfo, err error)
}

func NewDir(name string, reader func(count int) (items []os.FileInfo, err error)) *Dir {
	return &Dir{name, reader}
}

func (f *Dir) Close() error {
	return nil
}
func (f *Dir) Read(p []byte) (n int, err error) {
	return 0, ErrIsDir
}
func (f *Dir) Seek(offset int64, whence int) (int64, error) {
	return 0, ErrIsDir
}
func (f *Dir) Readdir(count int) (items []os.FileInfo, err error) {
	return f.Reader(count)
}
func (f *Dir) Stat() (os.FileInfo, error) {
	return oscommon.NewVirtualDirFileInfo(f.Name), nil
}
