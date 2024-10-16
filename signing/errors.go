package signing

import (
	"github.com/pritunl/tools/errors"
)

type LookupError struct {
	errors.DropboxError
}
