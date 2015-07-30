package repo

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Repo struct {
	Root string
}

func (r *Repo) Init() (err error) {
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
		path := filepath.Join(r.Root, dir)
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

func (r *Repo) getTargets() (targets []os.FileInfo, err error) {
	targets, err = ioutil.ReadDir(r.Root)
	if err != nil {
		err = &FileError{
			errors.Wrapf(err, "repo: Failed to read dir '%s'", r.Root),
		}
		return
	}

	return
}

func (r *Repo) createRedhat(distro, release, path string) (err error) {
	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		path+":/pacur", constants.DockerOrg+distro+"-"+release, "create",
		distro+"-"+release)
	if err != nil {
		return
	}

	repoPath := filepath.Join(r.Root, "mirror", "yum", distro, release)

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

func (r *Repo) createDebian(distro, release, path string) (err error) {
	confDir := filepath.Join(r.Root, distro+"-"+release, "conf")
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
		filepath.Join(r.Root, "mirror", "apt"))
	if err != nil {
		return
	}

	return
}

func (r *Repo) createTarget(target, path string) (err error) {
	distro, release, err := getDistro(target)
	if err != nil {
		return
	}

	switch distro {
	case "centos":
		err = r.createRedhat(distro, release, path)
	case "debian", "ubuntu":
		err = r.createDebian(distro, release, path)
	default:
		err = &UnknownType{
			errors.Newf("repo: Unknown repo type '%s'", target),
		}
	}

	return
}

func (r *Repo) Pull() (err error) {
	targets, err := r.getTargets()
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

func (r *Repo) Build() (err error) {
	targets, err := r.getTargets()
	if err != nil {
		return
	}

	for _, target := range targets {
		image := target.Name()
		if image == "mirror" || !target.IsDir() {
			continue
		}
		path := filepath.Join(r.Root, image)

		err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
			path+":/pacur", constants.DockerOrg+image)
		if err != nil {
			return
		}
	}

	return
}

func (r *Repo) Create() (err error) {
	targets, err := r.getTargets()
	if err != nil {
		return
	}

	for _, target := range targets {
		image := target.Name()
		if image == "mirror" || !target.IsDir() {
			continue
		}
		path := filepath.Join(r.Root, image)

		err = r.createTarget(image, path)
		if err != nil {
			return
		}
	}

	return
}
