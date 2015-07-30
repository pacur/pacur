package utils

import (
	"github.com/dropbox/godropbox/errors"
)

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

type ExecError struct {
	errors.DropboxError
}
