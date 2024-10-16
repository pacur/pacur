package resolver

import (
	"github.com/pritunl/tools/errors"
)

type ResolveError struct {
	errors.DropboxError
}
