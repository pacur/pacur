package mirror

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type Mirror struct {
	Root string
}

func (m *Mirror) createDebian(release string) (err error) {
	outDir := filepath.Join(m.Root, "apt")

	err = utils.MkdirAll(outDir)
	if err != nil {
		return
	}

	debs, err := utils.FindExt(m.Root, ".deb")
	if err != nil {
		return
	}

	for _, deb := range debs {
		err = utils.Exec(m.Root, "createrepo", "--outdir", outDir,
			"includedeb", release, deb)
		if err != nil {
			return
		}
	}

	return
}

func (m *Mirror) createRedhat() (err error) {
	err = utils.Exec(m.Root, "createrepo", ".")
	if err != nil {
		return
	}

	return
}

func (m *Mirror) Create(distro, release string) (err error) {
	switch distro {
	case "centos":
		err = m.createRedhat()
	case "debian":
		err = m.createDebian(release)
	default:
		err = &UnknownType{
			errors.Newf("mirror: Unknown type '%s'", distro),
		}
	}

	return
}
