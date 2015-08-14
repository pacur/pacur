package debian

import (
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type DebianProject struct {
	Root    string
	Path    string
	Distro  string
	Release string
}

func (p *DebianProject) Prep() (err error) {
	return
}

func (p *DebianProject) Create() (err error) {
	aptDir := filepath.Join(p.Path, "apt")

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		p.Path+":/pacur", constants.DockerOrg+p.Distro+"-"+p.Release,
		"create", p.Distro+"-"+p.Release)
	if err != nil {
		return
	}

	err = utils.Rsync(aptDir, filepath.Join(p.Root, "mirror", "apt"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(aptDir)
	if err != nil {
		return
	}

	err = utils.RemoveAll(filepath.Join(p.Path, "conf"))
	if err != nil {
		return
	}

	err = utils.RemoveAll(filepath.Join(p.Path, "db"))
	if err != nil {
		return
	}

	return
}
