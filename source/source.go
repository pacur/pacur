package source

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	path = 0
	url  = 1
)

type Source struct {
	Path   string
	Output string
}

func (s *Source) getType() int {
	if strings.HasPrefix(s.Path, "http") {
		return url
	}
	return path
}

func (s *Source) getUrl() (err error) {
	name, err := utils.HttpGet(s.Path, s.Output)
	if err != nil {
		return
	}

	cmd := exec.Command("tar", "xfz", filepath.Join(s.Output, name))
	cmd.Dir = s.Output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &GetError{
			errors.Wrapf(err, "builder: Failed to extract source '%s'",
				s.Path),
		}
		return
	}

	return
}

func (s *Source) getPath() (err error) {
	return
}

func (s *Source) Get() (err error) {
	switch s.getType() {
	case url:
		err = s.getUrl()
	case path:
		err = s.getPath()
	default:
		panic("utils: Unknown type")
	}
	if err != nil {
		return
	}

	return
}
