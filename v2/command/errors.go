package command

import (
	"github.com/dropbox/godropbox/errors"
)

type FileError struct {
	errors.DropboxError
}

type InvalidCommand struct {
	errors.DropboxError
}

type UnknownCommand struct {
	errors.DropboxError
}
