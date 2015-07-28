package redhat

import (
	"github.com/dropbox/godropbox/errors"
)

type WriteError struct {
	errors.DropboxError
}

type BuildError struct {
	errors.DropboxError
}
