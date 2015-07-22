package source

import (
	"github.com/dropbox/godropbox/errors"
)

type GetError struct {
	errors.DropboxError
}

type HashError struct {
	errors.DropboxError
}
