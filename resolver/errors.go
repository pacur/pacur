package resolver

import (
	"github.com/dropbox/godropbox/errors"
)

type ResolveError struct {
	errors.DropboxError
}
