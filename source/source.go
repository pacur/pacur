package source

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pacur/pacur/utils"
)

const (
	path = 0
	url  = 1
)

type Source struct {
	Root   string
	Hash   string
	Source string
	Output string
	Path   string
}

func (s *Source) getType() int {
	if strings.HasPrefix(s.Source, "http") {
		return url
	}
	return path
}

func (s *Source) parsePath() {
	s.Path = filepath.Join(s.Output, utils.Filename(s.Source))
}

func (s *Source) getUrl() (err error) {
	exists, err := utils.Exists(s.Path)
	if err != nil {
		return
	}

	if !exists {
		err = utils.HttpGet(s.Source, s.Path)
		if err != nil {
			return
		}
	}

	return
}

func (s *Source) getPath() (err error) {
	err = utils.Copy(s.Root, s.Source, s.Path, true)
	if err != nil {
		return
	}

	return
}

func (s *Source) extract() (err error) {
	var cmd *exec.Cmd

	if strings.HasSuffix(s.Path, ".tar") {
		cmd = exec.Command("tar", "--no-same-owner", "-xf", s.Path)
	} else if strings.HasSuffix(s.Path, ".zip") {
		cmd = exec.Command("unzip", s.Path)
	} else {
		split := strings.Split(s.Path, ".")
		if len(split) > 2 && split[len(split)-2] == "tar" {
			cmd = exec.Command("tar", "--no-same-owner", "-xf", s.Path)
		} else {
			return
		}
	}

	cmd.Dir = s.Output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return
	}

	return
}

func (s *Source) validate() (err error) {
	if strings.ToLower(s.Hash) == "skip" {
		return
	}

	file, err := os.Open(s.Path)
	if err != nil {
		return
	}
	defer file.Close()

	var hash hash.Hash
	switch len(s.Hash) {
	case 32:
		hash = md5.New()
	case 40:
		hash = sha1.New()
	case 64:
		hash = sha256.New()
	case 128:
		hash = sha512.New()
	default:
		return
	}

	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}

	sum := hash.Sum([]byte{})

	hexSum := fmt.Sprintf("%x", sum)

	if hexSum != s.Hash {
		return
	}

	return
}

func (s *Source) Get() (err error) {
	s.parsePath()

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

	err = s.validate()
	if err != nil {
		return
	}

	err = s.extract()
	if err != nil {
		return
	}

	return
}
