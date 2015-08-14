package redhat

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type RedhatProject struct {
	Root    string
	Path    string
	Distro  string
	Release string
}

func (p *RedhatProject) Prep() (err error) {
	return
}

func (p *RedhatProject) Create() (err error) {
	yumDir := filepath.Join(p.Path, "yum")

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		p.Path+":/pacur", constants.DockerOrg+p.Distro+"-"+p.Release,
		"create", p.Distro+"-"+p.Release)
	if err != nil {
		return
	}

	err = utils.Rsync(yumDir, filepath.Join(p.Root, "mirror", "yum"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(yumDir)
	if err != nil {
		return
	}

	return
}
