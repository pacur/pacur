package mirror

import (
	"fmt"
	"path/filepath"

	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/signing"
	"github.com/pacur/pacur/utils"
)

type Mirror struct {
	Name    string
	Root    string
	Distro  string
	Release string
	Signing bool
}

func (m *Mirror) createPacman() (err error) {
	outDir := filepath.Join(m.Root, "pacman")

	err = utils.MkdirAll(outDir)
	if err != nil {
		return
	}

	err = utils.RsyncExt(m.Root, outDir, ".pkg.tar.xz")
	if err != nil {
		return
	}

	if m.Signing {
		err = signing.SignPacman(outDir)
		if err != nil {
			return
		}
	}

	pkgs, err := utils.FindExt(outDir, ".pkg.tar.xz")
	for _, pkg := range pkgs {
		err = utils.Exec(outDir, "repo-add",
			fmt.Sprintf("%s.db.tar.gz", m.Name), pkg)
		if err != nil {
			return
		}
	}

	return
}

func (m *Mirror) createDebian() (err error) {
	outDir := filepath.Join(m.Root, "apt")
	confDir := filepath.Join(m.Root, "conf")
	confPath := filepath.Join(confDir, "distributions")

	err = utils.MkdirAll(confDir)
	if err != nil {
		return
	}

	data := "Origin: " + m.Name + "\nCodename: " + m.Release + "\n" +
		"Components: main\nArchitectures: i386 amd64\n"

	if m.Signing {
		id, e := signing.GetId()
		if e != nil {
			err = e
			return
		}

		data += fmt.Sprintf("SignWith: %s\n", id)

		err = utils.CreateWrite(filepath.Join(confDir, "options"),
			"ask-passphrase\n")
		if err != nil {
			return
		}
	}

	err = utils.CreateWrite(confPath, data)
	if err != nil {
		return
	}

	err = utils.MkdirAll(outDir)
	if err != nil {
		return
	}

	match, ok := constants.ReleasesMatch[m.Distro+"-"+m.Release]
	if !ok {
		return
	}

	debs, err := utils.FindMatch(m.Root, match)
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
	outDir := filepath.Join(m.Root, "yum", m.Distro, m.Release)

	err = utils.MkdirAll(outDir)
	if err != nil {
		return
	}

	match, ok := constants.ReleasesMatch[m.Distro+"-"+m.Release]
	if !ok {
		return
	}

	err = utils.RsyncMatch(m.Root, outDir, match)
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

		err = signing.CreateRedhatConf()
		if err != nil {
			return
		}
	}

	switch constants.DistroPack[m.Distro] {
	case "pacman":
		err = m.createPacman()
	case "debian":
		err = m.createDebian()
	case "redhat":
		err = m.createRedhat()
	default:
	}

	return
}
