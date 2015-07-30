package utils

import (
	"github.com/dropbox/godropbox/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func Chmod(path string, perm os.FileMode) (err error) {
	err = os.Chmod(path, perm)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err, "utils: Failed to chmod '%s'", path),
		}
		return
	}

	return
}

func Remove(path string) (err error) {
	err = os.Remove(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err, "utils: Failed to remove '%s'", path),
		}
		return
	}

	return
}

func RemoveAll(path string) (err error) {
	err = os.RemoveAll(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err, "utils: Failed to remove '%s'", path),
		}
		return
	}

	return
}

func ExistsMakeDir(path string) (err error) {
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = MkdirAll(path)
			if err != nil {
				return
			}
		} else {
			err = &MakeDirError{
				errors.Wrapf(err, "utils: Failed to stat '%s'", path),
			}
		}
		return
	}

	return
}

func Create(path string) (file *os.File, err error) {
	file, err = os.Create(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err, "utils: Failed to create '%s'", path),
		}
		return
	}

	return
}

func CreateWrite(path string, data string) (err error) {
	file, err := Create(path)
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

func Move(source, dest string) (err error) {
	err = Exec("", "mv", source, dest)
	if err != nil {
		return
	}

	return
}

func Copy(dir, source, dest string, presv bool) (err error) {
	args := []string{"-r", "-T", "-f"}

	if presv {
		args = append(args, "-p")
	}
	args = append(args, source, dest)

	err = Exec(dir, "cp", args...)
	if err != nil {
		return
	}

	return
}

func CopyFile(dir, source, dest string, presv bool) (err error) {
	args := []string{"-f"}

	if presv {
		args = append(args, "-p")
	}
	args = append(args, source, dest)

	err = Exec(dir, "cp", args...)
	if err != nil {
		return
	}

	return
}

func CopyFiles(source, dest string, presv bool) (err error) {
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

		err = CopyFile("", filepath.Join(source, file.Name()), dest, presv)
		if err != nil {
			return
		}
	}

	return
}

func FindExt(path, ext string) (matches []string, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		err = &ReadError{
			errors.Wrapf(err, "utils: Failed to read dir '%s'", path),
		}
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ext) {
			matches = append(matches, filepath.Join(path, file.Name()))
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

func GetDirSize(path string) (size int, err error) {
	output, err := ExecOutput("", "du", "-c", "-s", path)
	if err != nil {
		return
	}

	split := strings.Fields(output)

	size, err = strconv.Atoi(split[len(split)-2])
	if err != nil {
		err = &ReadError{
			errors.Wrapf(err, "utils: Failed to get dir size '%s'", path),
		}
		return
	}

	return
}

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
