package builder

import (
	"github.com/dropbox/godropbox/errors"
)

type BuilderError struct {
	errors.DropboxError
}
