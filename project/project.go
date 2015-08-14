package project

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/arch"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/debian"
	"github.com/pacur/pacur/redhat"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type Repo interface {
	Prep() error
	Create() error
}

type Project struct {
	Root string
}

func (p *Project) Init() (err error) {
	err = utils.MkdirAll(filepath.Join(p.Root, "mirror"))
	if err != nil {
		return
	}

	for _, release := range constants.Releases {
		err = utils.MkdirAll(filepath.Join("pkgname", release))
		if err != nil {
			return
		}
	}

	return
}

func (p *Project) getRepo(target, path string) (repo Repo, err error) {
	distro, release := getDistro(target)

	switch distro {
	case "archlinux":
		repo = &arch.ArchRepo{
			Root:    p.Root,
			Path:    path,
			Distro:  distro,
			Release: release,
		}
	case "centos":
		repo = &redhat.RedhatRepo{
			Root:    p.Root,
			Path:    path,
			Distro:  distro,
			Release: release,
		}
	case "debian", "ubuntu":
		repo = &debian.DebianRepo{
			Root:    p.Root,
			Path:    path,
			Distro:  distro,
			Release: release,
		}
	default:
		err = &UnknownType{
			errors.Newf("repo: Unknown repo type '%s'", target),
		}
	}

	return
}

func (p *Project) Pull() (err error) {
	for _, release := range constants.Releases {
		err = utils.Exec("", "docker", "pull", constants.DockerOrg+release)
		if err != nil {
			return
		}
	}

	return
}

func (p *Project) iterPackages(handle func(string, string) error) (err error) {
	projects, err := utils.ReadDir(p.Root)
	if err != nil {
		return
	}

	for _, project := range projects {
		if project.Name() == "mirror" || !project.IsDir() {
			continue
		}

		projectPath := filepath.Join(p.Root, project.Name())

		packages, e := utils.ReadDir(projectPath)
		if e != nil {
			err = e
			return
		}

		for _, pkg := range packages {
			err = handle(pkg.Name(), filepath.Join(projectPath, pkg.Name()))
			if err != nil {
				return
			}
		}
	}

	return
}

func (p *Project) Build() (err error) {
	err = p.iterPackages(func(release, path string) (err error) {
		err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
			path+":/pacur", constants.DockerOrg+release)
		if err != nil {
			return
		}

		return
	})
	if err != nil {
		return
	}

	return
}

func (p *Project) Repo() (err error) {
	return
}
