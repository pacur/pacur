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

type DistroProject interface {
	Prep() error
	Create() error
}

type Project struct {
	Root       string
	MirrorRoot string
	BuildRoot  string
}

func (p *Project) Init() {
	p.MirrorRoot = filepath.Join(p.Root, "mirror")
	p.BuildRoot = filepath.Join(p.Root, "mirror.tmp")
}

func (p *Project) InitProject() (err error) {
	err = utils.MkdirAll(p.MirrorRoot)
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

func (p *Project) getProject(target, path string) (
	proj DistroProject, err error) {

	distro, release := getDistro(target)

	switch distro {
	case "archlinux":
		proj = &arch.ArchProject{
			Root:       p.Root,
			MirrorRoot: p.MirrorRoot,
			BuildRoot:  p.BuildRoot,
			Path:       path,
			Distro:     distro,
			Release:    release,
		}
	case "centos":
		proj = &redhat.RedhatProject{
			Root:       p.Root,
			MirrorRoot: p.MirrorRoot,
			BuildRoot:  p.BuildRoot,
			Path:       path,
			Distro:     distro,
			Release:    release,
		}
	case "debian", "ubuntu":
		proj = &debian.DebianProject{
			Root:       p.Root,
			MirrorRoot: p.MirrorRoot,
			BuildRoot:  p.BuildRoot,
			Path:       path,
			Distro:     distro,
			Release:    release,
		}
	default:
		err = &UnknownType{
			errors.Newf("project: Unknown repo type '%s'", target),
		}
	}

	return
}

func (p *Project) iterPackages(filter string,
	handle func(string, string) error) (err error) {

	projects, err := utils.ReadDir(p.Root)
	if err != nil {
		return
	}

	for _, project := range projects {
		if project.Name() == "mirror" || !project.IsDir() {
			continue
		}

		if filter != "" && project.Name() != filter {
			continue
		}

		projectPath := filepath.Join(p.Root, project.Name())

		packages, e := utils.ReadDir(projectPath)
		if e != nil {
			err = e
			return
		}

		for _, pkg := range packages {
			if !pkg.IsDir() {
				continue
			}

			err = handle(pkg.Name(), filepath.Join(projectPath, pkg.Name()))
			if err != nil {
				return
			}
		}
	}

	return
}

func (p *Project) Build(filter string) (err error) {
	err = p.iterPackages(filter, func(target, path string) (err error) {
		err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
			path+":/pacur", constants.DockerOrg+target)
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

func (p *Project) Repo(filter string) (err error) {
	err = p.iterPackages(filter, func(target, path string) (err error) {
		proj, err := p.getProject(target, path)
		if err != nil {
			return
		}

		err = proj.Prep()
		if err != nil {
			return
		}

		return
	})
	if err != nil {
		return
	}

	for _, release := range constants.Releases {
		path := filepath.Join(p.BuildRoot, release)

		exists, e := utils.Exists(path)
		if e != nil {
			err = e
			return
		}

		if exists {
			proj, e := p.getProject(release, path)
			if e != nil {
				err = e
				return
			}

			err = proj.Create()
			if err != nil {
				return
			}
		}
	}

	err = utils.RemoveAll(p.BuildRoot)
	if err != nil {
		return
	}

	return
}
