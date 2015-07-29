package repo

import (
	"github.com/dropbox/godropbox/errors"
)

type FileError struct {
	errors.DropboxError
}

type BuildError struct {
	errors.DropboxError
}

type UnknownType struct {
	errors.DropboxError
}
