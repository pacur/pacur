package project

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Project struct {
	Root string
}

func (p *Project) Init() (err error) {
	for _, dir := range []string{
		"mirror/yum/centos/6",
		"mirror/yum/centos/7",
		"mirror/apt/debian/dists/jessie",
		"mirror/apt/debian/dists/wheezy",
		"mirror/apt/ubuntu/dists/precise",
		"mirror/apt/ubuntu/dists/trusty",
		"mirror/apt/ubuntu/dists/vivid",
		"mirror/apt/ubuntu/dists/wily",
		"centos-6",
		"centos-7",
		"debian-jessie",
		"debian-wheezy",
		"ubuntu-precise",
		"ubuntu-trusty",
		"ubuntu-vivid",
		"ubuntu-wily",
	} {
		path := filepath.Join(p.Root, dir)
		err = os.MkdirAll(path, 0755)
		if err != nil {
			err = &FileError{
				errors.Wrapf(err, "repo: Failed to mkdir '%s'", path),
			}
			return
		}
	}

	return
}

func (p *Project) getTargets() (targets []os.FileInfo, err error) {
	targets, err = ioutil.ReadDir(p.Root)
	if err != nil {
		err = &FileError{
			errors.Wrapf(err, "repo: Failed to read dir '%s'", p.Root),
		}
		return
	}

	return
}

func (p *Project) createRedhat(distro, release, path string) (err error) {
	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		path+":/pacur", constants.DockerOrg+distro+"-"+release, "create",
		distro+"-"+release)
	if err != nil {
		return
	}

	repoPath := filepath.Join(p.Root, "mirror", "yum", distro, release)

	err = utils.RsyncExt(path, repoPath, "rpm")
	if err != nil {
		return
	}

	err = utils.Rsync(filepath.Join(path, "repodata"),
		filepath.Join(repoPath, "repodata"))
	if err != nil {
		return
	}

	return
}

func (p *Project) createDebian(distro, release, path string) (err error) {
	confDir := filepath.Join(p.Root, distro+"-"+release, "conf")
	confPath := filepath.Join("conf", "distributions")

	err = utils.MkdirAll(confDir)
	if err != nil {
		return
	}

	err = utils.CreateWrite(confPath, "Codename: "+release+"\n"+
		"Components: main\nArchitectures: amd64\n")
	if err != nil {
		return
	}

	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		path+":/pacur", constants.DockerOrg+distro+"-"+release, "create",
		distro+"-"+release)
	if err != nil {
		return
	}

	err = utils.Rsync(filepath.Join(path, "apt"),
		filepath.Join(p.Root, "mirror", "apt"))
	if err != nil {
		return
	}

	return
}

func (p *Project) createTarget(target, path string) (err error) {
	distro, release, err := getDistro(target)
	if err != nil {
		return
	}

	switch distro {
	case "centos":
		err = p.createRedhat(distro, release, path)
	case "debian", "ubuntu":
		err = p.createDebian(distro, release, path)
	default:
		err = &UnknownType{
			errors.Newf("repo: Unknown repo type '%s'", target),
		}
	}

	return
}

func (p *Project) Pull() (err error) {
	targets, err := p.getTargets()
	if err != nil {
		return
	}

	for _, target := range targets {
		image := target.Name()
		if image == "mirror" || !target.IsDir() {
			continue
		}

		err = utils.Exec("", "docker", "pull", constants.DockerOrg+image)
		if err != nil {
			return
		}
	}

	return
}

func (p *Project) Build() (err error) {
	targets, err := p.getTargets()
	if err != nil {
		return
	}

	for _, target := range targets {
		image := target.Name()
		if image == "mirror" || !target.IsDir() {
			continue
		}
		path := filepath.Join(p.Root, image)

		err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
			path+":/pacur", constants.DockerOrg+image)
		if err != nil {
			return
		}
	}

	return
}

func (p *Project) Repo() (err error) {
	targets, err := p.getTargets()
	if err != nil {
		return
	}

	for _, target := range targets {
		image := target.Name()
		if image == "mirror" || !target.IsDir() {
			continue
		}
		path := filepath.Join(p.Root, image)

		err = p.createTarget(image, path)
		if err != nil {
			return
		}
	}

	return
}
