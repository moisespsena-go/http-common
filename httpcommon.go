package httpcommon

import (
	"os"

	"github.com/moisespsena-go/io-common"
)

type File struct {
	FileInfo os.FileInfo
	open     func() (iocommon.ReadSeekCloser, error)
	seeker   iocommon.ReadSeekCloser
}

func NewFile(info os.FileInfo, open func() (iocommon.ReadSeekCloser, error)) *File {
	return &File{info, open, nil}
}

func (f *File) Close() error {
	if f.seeker != nil {
		return f.seeker.Close()
	}
	return nil
}

func (f *File) Read(p []byte) (n int, err error) {
	if f.seeker == nil {
		if f.seeker, err = f.open(); err != nil {
			return
		}
	}
	return f.seeker.Read(p)
}

func (f *File) Seek(offset int64, whence int) (n int64, err error) {
	if f.seeker == nil {
		if f.seeker, err = f.open(); err != nil {
			return
		}
	}
	return f.seeker.Seek(offset, whence)
}

func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	return nil, ErrIsFile
}

func (f *File) Stat() (os.FileInfo, error) {
	return f.FileInfo, nil
}
