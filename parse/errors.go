package parse

import (
	"github.com/dropbox/godropbox/errors"
)

type SyntaxError struct {
	errors.DropboxError
}
