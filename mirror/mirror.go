package mirror

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type Mirror struct {
	Root    string
	Distro  string
	Release string
}

func (m *Mirror) createDebian() (err error) {
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
			"includedeb", m.Release, deb)
		if err != nil {
			return
		}
	}

	return
}

func (m *Mirror) createRedhat(release string) (err error) {
	outDir := filepath.Join(m.Root, "yum", "centos", m.Release)

	err = utils.RsyncExt(m.Root, outDir, ".rpm")
	if err != nil {
		return
	}

	err = utils.Exec(outDir, "createrepo", ".")
	if err != nil {
		return
	}

	return
}

func (m *Mirror) Create() (err error) {
	switch distro {
	case "centos":
		err = m.createRedhat()
	case "debian", "ubuntu":
		err = m.createDebian(release)
	default:
		err = &UnknownType{
			errors.Newf("mirror: Unknown type '%s'", distro),
		}
	}

	return
}
