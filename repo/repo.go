package repo

import (
	"github.com/dropbox/godropbox/errors"
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
		"mirror",
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

func (r *Repo) Build() (err error) {
	targets, err := ioutil.ReadDir(r.Root)
	if err != nil {
		err = &FileError{
			errors.Wrapf(err, "repo: Failed to read dir '%s'", r.Root),
		}
		return
	}

	for _, target := range targets {
		name := target.Name()

		if name == "mirror" || !target.IsDir() {
			continue
		}
		path := filepath.Join(r.Root, name)

		cmd := exec.Command("docker", "pull", "pacur/"+name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			err = &BuildError{
				errors.Wrapf(err, "repo: Failed to pull 'pacur/%s'", name),
			}
			return
		}

		cmd = exec.Command("docker", "run", "--rm", "-t", "-v",
			path+":/pacur", "pacur/"+name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			err = &BuildError{
				errors.Wrapf(err, "repo: Failed to build '%s'", path),
			}
			return
		}
	}

	return
}

func (r *Repo) createRedhat() (err error) {
	cmd := exec.Command("createrepo", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.Root

	err = cmd.Run()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "repo: Failed to create redhat repo '%s'",
				r.Root),
		}
		return
	}

	return
}

func (r *Repo) Create(typ string) (err error) {
	switch typ {
	case "redhat":
		err = r.createRedhat()
	default:
		err = &UnknownType{
			errors.Newf("repo: Unknown type '%s'", typ),
		}
	}

	return
}
