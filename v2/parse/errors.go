package parse

import (
	"github.com/dropbox/godropbox/errors"
)

type SyntaxError struct {
	errors.DropboxError
}

type FileError struct {
	errors.DropboxError
}
