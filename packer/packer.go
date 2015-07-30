package packer

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/debian"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/redhat"
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
	case "centos":
		pcker = &redhat.Redhat{
			Pack:    pac,
			Distro:  distro,
			Release: release,
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
