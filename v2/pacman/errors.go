package pacman

import (
	"github.com/dropbox/godropbox/errors"
)

type BuildError struct {
	errors.DropboxError
}
