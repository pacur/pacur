package pack

import (
	"github.com/dropbox/godropbox/errors"
)

type ParseError struct {
	errors.DropboxError
}

type ValidationError struct {
	errors.DropboxError
}
