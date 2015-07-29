package repo

import (
	"github.com/dropbox/godropbox/errors"
	"strings"
)

func GetRepoType(name string) (typ string, err error) {
	name = strings.Split(name, "-")[0]

	switch name {
	case "centos":
		typ = "redhat"
	case "debian", "ubuntu":
		typ = "debian"
	default:
		err = &UnknownType{
			errors.Newf("repo: Unknown repo type '%s'", name),
		}
	}

	return
}
