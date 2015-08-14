package arch

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type ArchProject struct {
	Root    string
	Path    string
	Distro  string
	Release string
}

func (p *ArchProject) Prep() (err error) {
	return
}

func (p *ArchProject) Create() (err error) {
	archDir := filepath.Join(p.Path, "arch")

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		p.Path+":/pacur", constants.DockerOrg+p.Distro, "create",
		p.Distro)
	if err != nil {
		return
	}

	err = utils.Rsync(archDir, filepath.Join(p.Root, "mirror", "arch"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(archDir)
	if err != nil {
		return
	}

	return
}
