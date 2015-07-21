package source

import (
	"github.com/pacur/pacur/utils"
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
	err = utils.HttpGet(s.Path, s.Output)
	if err != nil {
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
