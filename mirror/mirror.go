package mirror

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/signing"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type Mirror struct {
	Root    string
	Distro  string
	Release string
	Signing bool
}

func (m *Mirror) createArch() (err error) {
	outDir := filepath.Join(m.Root, "arch")

	err = utils.MkdirAll(outDir)
	if err != nil {
		return
	}

	err = utils.RsyncExt(m.Root, outDir, ".pkg.tar.xz")
	if err != nil {
		return
	}

	pkgs, err := utils.FindExt(outDir, ".pkg.tar.xz")
	for _, pkg := range pkgs {
		err = utils.Exec(outDir, "repo-add", "pacur.db.tar.gz", pkg)
		if err != nil {
			return
		}
	}

	return
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
		err = utils.Exec(m.Root, "reprepro", "--outdir", outDir,
			"includedeb", m.Release, deb)
		if err != nil {
			return
		}
	}

	return
}

func (m *Mirror) createRedhat() (err error) {
	outDir := filepath.Join(m.Root, "yum", "centos", m.Release)

	err = utils.MkdirAll(outDir)
	if err != nil {
		return
	}

	err = utils.RsyncExt(m.Root, outDir, ".rpm")
	if err != nil {
		return
	}

	if m.Signing {
		err = signing.SignRedhat(outDir)
		if err != nil {
			return
		}
	}

	err = utils.Exec(outDir, "createrepo", ".")
	if err != nil {
		return
	}

	return
}

func (m *Mirror) Create() (err error) {
	keyPath := filepath.Join(m.Root, "sign.key")

	m.Signing, err = utils.Exists(keyPath)
	if err != nil {
		return
	}

	if m.Signing {
		err = signing.ImportKey(keyPath)
		if err != nil {
			return
		}
	}

	switch m.Distro {
	case "archlinux":
		err = m.createArch()
	case "centos":
		err = m.createRedhat()
	case "debian", "ubuntu":
		err = m.createDebian()
	default:
		err = &UnknownType{
			errors.Newf("mirror: Unknown type '%s'", m.Distro),
		}
	}

	return
}
