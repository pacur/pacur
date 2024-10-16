package source

import (
	"github.com/pritunl/tools/errors"
)

type GetError struct {
	errors.DropboxError
}

type HashError struct {
	errors.DropboxError
}
