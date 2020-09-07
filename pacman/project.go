package pacman

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type PacmanProject struct {
	Name       string
	Root       string
	MirrorRoot string
	BuildRoot  string
	Path       string
	Distro     string
	Release    string
}

func (p *PacmanProject) getBuildDir() (path string, err error) {
	path = filepath.Join(p.BuildRoot, p.Distro)

	err = utils.MkdirAll(path)
	if err != nil {
		return
	}

	return
}

func (p *PacmanProject) Prep() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	keyPath := filepath.Join(p.Path, "..", "sign.key")
	exists, err := utils.Exists(keyPath)
	if err != nil {
		return
	}

	if exists {
		err = utils.CopyFile("", keyPath, buildDir, true)
		if err != nil {
			return
		}
	}

	err = utils.RsyncExt(p.Path, buildDir, ".pkg.tar.zst")
	if err != nil {
		return
	}

	return
}

func (p *PacmanProject) Create() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		buildDir+":/pacur:Z", constants.DockerOrg+p.Distro, "create",
		p.Distro, p.Name)
	if err != nil {
		return
	}

	err = utils.Rsync(filepath.Join(buildDir, "pacman"),
		filepath.Join(p.MirrorRoot, "pacman"))
	if err != nil {
		return
	}

	return
}
