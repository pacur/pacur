package project

import (
	"github.com/pritunl/tools/errors"
)

type FileError struct {
	errors.DropboxError
}

type ParseError struct {
	errors.DropboxError
}

type UnknownType struct {
	errors.DropboxError
}
