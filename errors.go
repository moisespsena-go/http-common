package httpcommon

import "errors"

var (
	ErrIsDir  = errors.New("is directory")
	ErrIsFile = errors.New("is file")
)
