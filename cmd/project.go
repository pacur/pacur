package cmd

import (
	"flag"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/project"
	"os"
)

func Project() (err error) {
	path, err := os.Getwd()
	if err != nil {
		err = &FileError{
			errors.Wrapf(err, "cmd: Failed to get working directory"),
		}
		return
	}

	proj := &project.Project{
		Root: path,
	}

	cmd := flag.Arg(1)
	switch cmd {
	case "init":
		err = proj.Init()
	case "build":
		err = proj.Build()
	default:
		err = &UnknownCommand{
			errors.Newf("cmd: Unknown cmd '%s'", cmd),
		}
	}

	return
}
