package repo

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"io/ioutil"
	"os"
	"os/exec"
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

func (r *Repo) createRedhat(distro, release, path string) (err error) {
	err = utils.Exec("", "docker", "run", "--rm", "-t", "-v",
		path+":/pacur", "pacur/"+distro+"-"+release, "create",
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
	confDir := filepath.Join(r.Root, "conf")
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

	cmd := exec.Command("docker", "run", "--rm", "-t", "-v",
		path+":/pacur", "pacur/"+distro+"-"+release, "create",
		distro+"-"+release)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "repo: Failed to build '%s'", path),
		}
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

func (r *Repo) Create(image, path string) (err error) {
	distro, release, err := getDistro(image)
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
			errors.Newf("repo: Unknown repo type '%s'", image),
		}
	}

	return
}

func (r *Repo) Build() (err error) {
	targets, err := ioutil.ReadDir(r.Root)
	if err != nil {
		err = &FileError{
			errors.Wrapf(err, "repo: Failed to read dir '%s'", r.Root),
		}
		return
	}

	for _, target := range targets {
		image := target.Name()

		if image == "mirror" || !target.IsDir() {
			continue
		}
		path := filepath.Join(r.Root, image)

		cmd := exec.Command("docker", "pull", "pacur/"+image)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			err = &BuildError{
				errors.Wrapf(err, "repo: Failed to pull 'pacur/%s'", image),
			}
			return
		}

		cmd = exec.Command("docker", "run", "--rm", "-t", "-v",
			path+":/pacur", "pacur/"+image)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			err = &BuildError{
				errors.Wrapf(err, "repo: Failed to build '%s'", path),
			}
			return
		}

		err = r.Create(image, path)
		if err != nil {
			return
		}
	}

	return
}
