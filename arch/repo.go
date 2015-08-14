package arch

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type ArchRepo struct {
	Root    string
	Path    string
	Distro  string
	Release string
}

func (r *ArchRepo) Prep() (err error) {
	return
}

func (r *ArchRepo) Create() (err error) {
	archDir := filepath.Join(r.Path, "arch")

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		r.Path+":/pacur", constants.DockerOrg+r.Distro, "create",
		r.Distro)
	if err != nil {
		return
	}

	err = utils.Rsync(archDir, filepath.Join(r.Root, "mirror", "arch"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(archDir)
	if err != nil {
		return
	}

	return
}
