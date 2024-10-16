package pacman

import (
	"github.com/pritunl/tools/errors"
)

type BuildError struct {
	errors.DropboxError
}
