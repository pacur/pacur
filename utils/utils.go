package utils

import (
	"github.com/dropbox/godropbox/errors"
	"os"
)

func ExistsMakeDir(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, 755)
		if err != nil {
			err = &MakeDirError{
				errors.Wrap(err, "utils: Failed to stat dir"),
			}
			return
		}
	} else if err != nil {
		err = &MakeDirError{
			errors.Wrap(err, "utils: Failed to create dir"),
		}
		return
	}

	return
}
