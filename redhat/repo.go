package redhat

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type RedhatRepo struct {
	Root    string
	Path    string
	Distro  string
	Release string
}

func (r *RedhatRepo) Prep() (err error) {
	return
}

func (r *RedhatRepo) Create() (err error) {
	yumDir := filepath.Join(r.Path, "yum")

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		r.Path+":/pacur", constants.DockerOrg+r.Distro+"-"+r.Release,
		"create", r.Distro+"-"+r.Release)
	if err != nil {
		return
	}

	err = utils.Rsync(yumDir, filepath.Join(r.Root, "mirror", "yum"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(yumDir)
	if err != nil {
		return
	}

	return
}
