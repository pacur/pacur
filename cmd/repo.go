package cmd

import (
	"github.com/pacur/pacur/repo"
	"os"
	"github.com/dropbox/godropbox/errors"
	"flag"
)

func Repo() (err error) {
	path, err := os.Getwd()
	if err != nil {
		err = &FileError{
			errors.Wrapf(err, "cmd: Failed to get working directory"),
		}
		return
	}

	rpo := &repo.Repo{
		Root: path,
	}

	cmd := flag.Arg(1)
	switch cmd {
	case "init":
		err = rpo.Init()
	case "build":
		err = rpo.Build()
	default:
		err = &UnknownCommand{
			errors.Newf("cmd: Unknown cmd '%s'", cmd),
		}
	}

	return
}
