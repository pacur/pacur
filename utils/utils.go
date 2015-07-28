package utils

import (
	"github.com/dropbox/godropbox/errors"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"io/ioutil"
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

func GetDirSize(path string) (size int, err error) {
	cmd := exec.Command("du", "-c", "-s", path)
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		err = &ReadError{
			errors.Wrapf(err, "utils: Failed to get dir size '%s'", path),
		}
		return
	}

	split := strings.Fields(string(output))

	size, err = strconv.Atoi(split[len(split)-2])
	if err != nil {
		err = &ReadError{
			errors.Wrapf(err, "utils: Failed to get dir size '%s'", path),
		}
		return
	}

	return
}

func Move(source, dest string) (err error) {
	cmd := exec.Command("mv", source, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &CopyError{
			errors.Wrapf(err, "utils: Failed to move '%s' to '%s'",
				source, dest),
		}
		return
	}

	return
}

func Copy(dir, source, dest string) (err error) {
	cmd := exec.Command("cp", "-p", "-r", "-T", "-f", source, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir

	err = cmd.Run()
	if err != nil {
		err = &CopyError{
			errors.Wrapf(err, "utils: Failed to copy '%s' to '%s'",
				source, dest),
		}
		return
	}

	return
}

func CopyFile(source, dest string) (err error) {
	cmd := exec.Command("cp", "-f", source, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &CopyError{
			errors.Wrapf(err, "utils: Failed to copy '%s' to '%s'",
				source, dest),
		}
		return
	}

	return
}

func CopyFiles(source, dest string) (err error) {
	files, err := ioutil.ReadDir(source)
	if err != nil {
		err = &ReadError{
			errors.Wrapf(err, "utils: Failed to read dir '%s'", source),
		}
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		err = CopyFile(filepath.Join(source, file.Name()), dest)
		if err != nil {
			return
		}
	}

	return
}

func Filename(path string) string {
	n := strings.LastIndex(path, "/")
	if n == -1 {
		return path
	}

	return path[n+1:]
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
