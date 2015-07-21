package utils

import (
	"os"
)

func ExistsMakeDir(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, 755)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	}

	return
}
