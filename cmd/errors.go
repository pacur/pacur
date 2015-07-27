package cmd

import (
	"github.com/dropbox/godropbox/errors"
)

type UnknownCommand struct {
	errors.DropboxError
}
