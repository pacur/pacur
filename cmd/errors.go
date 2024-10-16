package cmd

import (
	"github.com/pritunl/tools/errors"
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
