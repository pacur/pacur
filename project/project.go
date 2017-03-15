package project

import (
	"encoding/json"
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/debian"
	"github.com/pacur/pacur/pacman"
	"github.com/pacur/pacur/redhat"
	"github.com/pacur/pacur/utils"
	"path/filepath"
	"strings"
)

type DistroProject interface {
	Prep() error
	Create() error
}

type conf struct {
	Name string `json:"name"`
}

type Project struct {
	confPath   string
	Root       string
	MirrorRoot string
	BuildRoot  string
	Name       string
}

func (p *Project) Init() (err error) {
	p.MirrorRoot = filepath.Join(p.Root, "mirror")
	p.BuildRoot = filepath.Join(p.Root, "mirror.tmp")
	p.confPath = filepath.Join(p.Root, "pacur.json")

	exists, err := utils.Exists(p.confPath)
	if err != nil {
		return
	}

	if exists {
		dataByt, e := utils.ReadFile(p.confPath)
		if e != nil {
			err = e
			return
		}

		data := conf{}

		err = json.Unmarshal(dataByt, &data)
		if err != nil {
			err = &ParseError{
				errors.Wrapf(err,
					"project: Failed to parse project conf '%s'", p.confPath),
			}
			return
		}

		p.Name = data.Name
	} else {
		p.Name = "pacur"
	}

	return
}

func (p *Project) InitProject() (err error) {
	exists, err := utils.Exists(p.confPath)
	if err != nil {
		return
	}

	if !exists {
		err = utils.CreateWrite(p.confPath, `{\n    "name": "pacur"\n}\n`)
		if err != nil {
			return
		}
	}

	return
}

func (p *Project) getProject(target, path string) (
	proj DistroProject, err error) {

	distro, release := getDistro(target)

	switch constants.DistroPack[distro] {
	case "pacman":
		proj = &pacman.PacmanProject{
			Name:       p.Name,
			Root:       p.Root,
			MirrorRoot: p.MirrorRoot,
			BuildRoot:  p.BuildRoot,
			Path:       path,
			Distro:     distro,
			Release:    release,
		}
	case "debian":
		proj = &debian.DebianProject{
			Name:       p.Name,
			Root:       p.Root,
			MirrorRoot: p.MirrorRoot,
			BuildRoot:  p.BuildRoot,
			Path:       path,
			Distro:     distro,
			Release:    release,
		}
	case "redhat":
		proj = &redhat.RedhatProject{
			Name:       p.Name,
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
		if strings.HasPrefix(project.Name(), ".") ||
			project.Name() == "mirror" || !project.IsDir() {

			continue
		}

		if filter != "" && project.Name() != filter {
			continue
		}

		projectPath := filepath.Join(p.Root, project.Name())

		targets, e := getTargets(projectPath)
		if e != nil {
			err = e
			return
		}

		for _, target := range targets {
			err = handle(target, projectPath)
			if err != nil {
				return
			}
		}
	}

	return
}

func (p *Project) Build(filter string) (err error) {
	err = p.iterPackages(filter, func(target, path string) (err error) {
		fmt.Println("******************************************************")
		fmt.Printf("Building: %s - %s\n", filter, target)
		fmt.Println("******************************************************")

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
	_ = utils.RemoveAll(p.BuildRoot)
	_ = utils.RemoveAll(p.MirrorRoot)

	err = utils.MkdirAll(p.MirrorRoot)
	if err != nil {
		return
	}

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
