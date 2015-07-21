package utils

import (
	"github.com/dropbox/godropbox/errors"
)

type MakeDirError struct {
	errors.DropboxError
}
