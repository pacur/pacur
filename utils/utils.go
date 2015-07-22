package utils

import (
	"github.com/dropbox/godropbox/errors"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	chars = []rune(
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func UrlFilename(url string) string {
	n := strings.LastIndex(url, "/")
	if n == -1 {
		return ""
	}
	return url[n+1:]
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

func HttpGet(url, outputDir string) (name string, err error) {
	name = UrlFilename(url)
	if name == "" {
		err = &InvalidUrl{
			errors.Newf("utils: Failed to get filename from '%s'", url),
		}
	}
	output := filepath.Join(outputDir, name)

	cmd := exec.Command("wget", url, "-O", output)
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
