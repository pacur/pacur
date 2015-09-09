package packer

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/debian"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/pacman"
	"github.com/pacur/pacur/redhat"
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
			errors.Newf("packer: Unkown system %s", system),
		}
		return
	}

	return
}
