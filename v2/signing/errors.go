package signing

import (
	"github.com/dropbox/godropbox/errors"
)

type LookupError struct {
	errors.DropboxError
}
