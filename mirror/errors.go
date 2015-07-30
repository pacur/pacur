package mirror

import (
	"github.com/dropbox/godropbox/errors"
)

type UnknownType struct {
	errors.DropboxError
}
