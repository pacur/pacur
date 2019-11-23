package packer

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/m0rf30/pacur/constants"
	"github.com/m0rf30/pacur/debian"
	"github.com/m0rf30/pacur/pack"
	"github.com/m0rf30/pacur/pacman"
	"github.com/m0rf30/pacur/redhat"
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
