package debian

import (
	"path/filepath"

	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
)

type DebianProject struct {
	Name       string
	Root       string
	MirrorRoot string
	BuildRoot  string
	Path       string
	Distro     string
	Release    string
}

func (p *DebianProject) getBuildDir() (path string, err error) {
	path = filepath.Join(p.BuildRoot, p.Distro+"-"+p.Release)

	err = utils.MkdirAll(path)
	if err != nil {
		return
	}

	return
}

func (p *DebianProject) Prep() (err error) {
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

	err = utils.RsyncRelExt(p.Path, buildDir, p.Release, ".deb")
	if err != nil {
		return
	}

	return
}

func (p *DebianProject) Create() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	err = utils.Exec("", "podman", "run", "--rm", "-t", "-v",
		buildDir+":/pacur:Z", constants.DockerOrg+p.Distro+"-"+p.Release,
		"create", p.Distro+"-"+p.Release, p.Name)
	if err != nil {
		return
	}

	err = utils.Rsync(filepath.Join(buildDir, "apt"),
		filepath.Join(p.MirrorRoot, "apt"))
	if err != nil {
		return
	}

	return
}

func (p *DebianProject) Clean() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	err = utils.RemoveAll(buildDir)
	if err != nil {
		return
	}

	return
}
