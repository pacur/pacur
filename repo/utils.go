package repo

import (
	"github.com/dropbox/godropbox/errors"
	"strings"
)

func getDistro(name string) (distro, release string, err error) {
	split := strings.Split(name, "-")
	if len(split) < 2 {
		err = &UnknownType{
			errors.Newf("repo: Unknown distro '%s'", name),
		}
		return
	}
	distro = split[0]
	release = split[1]

	return
}
