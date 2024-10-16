package parse

import (
	"github.com/pritunl/tools/errors"
)

type SyntaxError struct {
	errors.DropboxError
}

type FileError struct {
	errors.DropboxError
}
