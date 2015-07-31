package arch

import (
	"github.com/dropbox/godropbox/errors"
)

type BuildError struct {
	errors.DropboxError
}
