package packer

import (
	"github.com/dropbox/godropbox/errors"
)

type UnknownSystem struct {
	errors.DropboxError
}
