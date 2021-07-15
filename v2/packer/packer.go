package packer

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/aanno/pacur/v2/constants"
	"github.com/aanno/pacur/v2/debian"
	"github.com/aanno/pacur/v2/pack"
	"github.com/aanno/pacur/v2/pacman"
	"github.com/aanno/pacur/v2/redhat"
)

type Packer interface {
	Prep() error
	Build() error
}

func GetPacker(pac *pack.Pack, distro, release string) (
	pcker Packer, err error) {

	switch constants.DistroPack[distro] {
	case "pacman":
		pcker = &pacman.Pacman{
			Pack: pac,
		}
	case "debian":
		pcker = &debian.Debian{
			Pack: pac,
		}
	case "redhat":
		pcker = &redhat.Redhat{
			Pack: pac,
		}
	default:
		system := distro
		if release != "" {
			system += "-" + release
		}

		err = &UnknownSystem{
			errors.Newf("packer: Unknown system %s", system),
		}
		return
	}

	return
}
