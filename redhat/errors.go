package redhat

import (
	"github.com/dropbox/godropbox/errors"
)

type HashError struct {
	errors.DropboxError
}

type WriteError struct {
	errors.DropboxError
}

type BuildError struct {
	errors.DropboxError
}
