package packer

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/debian"
	"github.com/pacur/pacur/pack"
)

type Packer interface {
	Prep() error
	Build() error
}

func GetPacker(pac *pack.Pack, distro, release string) (
	pcker Packer, err error) {

	switch distro {
	case "debian", "ubuntu":
		pcker = &debian.Debian{
			Pack:    pac,
			Distro:  distro,
			Release: release,
		}
	default:
		err = &UnknownSystem{
			errors.Newf("packer: Unkown system %s-%s", distro, release),
		}
		return
	}

	return
}
