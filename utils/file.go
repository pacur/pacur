package utils

import (
	"github.com/dropbox/godropbox/errors"
	"os"
)

func MkdirAll(path string) (err error) {
	err = os.MkdirAll(path, 0755)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err, "utils: Failed to mkdir '%s'", path),
		}
		return
	}

	return
}

func Create(path string, perm os.FileMode) (file *os.File, err error) {
	file, err = os.Create(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err, "utils: Failed to create '%s'", path),
		}
		return
	}

	return
}

func CreateWrite(path string, perm os.FileMode, data string) (err error) {
	file, err := Create(path, perm)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err, "utils: Failed to write to file '%s'", path),
		}
		return
	}

	return
}
