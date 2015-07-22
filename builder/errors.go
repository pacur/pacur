package builder

import (
	"github.com/dropbox/godropbox/errors"
)

type BuilderError struct {
	errors.DropboxError
}

type ScriptError struct {
	errors.DropboxError
}
