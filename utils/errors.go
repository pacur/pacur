package utils

import (
	"github.com/dropbox/godropbox/errors"
)

type MakeDirError struct {
	errors.DropboxError
}

type HttpError struct {
	errors.DropboxError
}

type ExistsError struct {
	errors.DropboxError
}

type CopyError struct {
	errors.DropboxError
}

type ReadError struct {
	errors.DropboxError
}

type WriteError struct {
	errors.DropboxError
}
