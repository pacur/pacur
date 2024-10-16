package mirror

import (
	"github.com/pritunl/tools/errors"
)

type BuildError struct {
	errors.DropboxError
}

type UnknownType struct {
	errors.DropboxError
}
