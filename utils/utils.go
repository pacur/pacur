package utils

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/m0rf30/pacur/constants"
	"math/rand"
	"os"
	"os/exec"
)

var (
	chars = []rune(
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

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

func PullContainers() (err error) {
	for _, release := range constants.Releases {
		err = Exec("", "docker", "pull", constants.DockerOrg+release)
		if err != nil {
			return
		}
	}

	return
}
