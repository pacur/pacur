package debian

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type DebianRepo struct {
	Root    string
	Path    string
	Distro  string
	Release string
}

func (r *DebianRepo) Prep() (err error) {
	return
}

func (r *DebianRepo) Create() (err error) {
	aptDir := filepath.Join(r.Path, "apt")

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		r.Path+":/pacur", constants.DockerOrg+r.Distro+"-"+r.Release,
		"create", r.Distro+"-"+r.Release)
	if err != nil {
		return
	}

	err = utils.Rsync(aptDir, filepath.Join(r.Root, "mirror", "apt"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(aptDir)
	if err != nil {
		return
	}

	err = utils.RemoveAll(filepath.Join(r.Path, "conf"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(filepath.Join(r.Path, "db"))
	if err != nil {
		return
	}

	return
}
