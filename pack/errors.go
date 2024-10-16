package pack

import (
	"github.com/pritunl/tools/errors"
)

type ParseError struct {
	errors.DropboxError
}

type ValidationError struct {
	errors.DropboxError
}
