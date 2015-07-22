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

type InvalidPath struct {
	errors.DropboxError
}

type ExistsError struct {
	errors.DropboxError
}
