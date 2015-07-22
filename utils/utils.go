package utils

import (
	"github.com/dropbox/godropbox/errors"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

var (
	chars = []rune(
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func Exists(path string) (exists bool, err error) {
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		} else {
			err = &ExistsError{
				errors.Wrapf(err, "utils: Exists check error for '%s'", path),
			}
		}
	} else {
		exists = true
	}

	return
}

func Filename(url string) (name string, err error) {
	n := strings.LastIndex(url, "/")
	if n == -1 {
		err = &InvalidPath{
			errors.Newf("utils: Failed to get filename from '%s'", url),
		}
		return
	}
	name = url[n+1:]

	return
}

func ExistsMakeDir(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
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

func HttpGet(url, output string) (err error) {
	cmd := exec.Command("wget", url, "-O", output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &HttpError{
			errors.Wrapf(err, "utils: Failed to get '%s'", url),
		}
		return
	}

	return
}

func RandStr(n int) (str string) {
	strList := make([]rune, n)
	for i := range strList {
		strList[i] = chars[rand.Intn(len(chars))]
	}
	str = string(strList)
	return
}
