package cmd

import (
	"flag"
	"os"

	"github.com/pacur/pacur/project"
	"github.com/pritunl/tools/errors"
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
	err = proj.Init()
	if err != nil {
		return
	}

	cmd := flag.Arg(1)
	switch cmd {
	case "init":
		err = proj.InitProject()
	case "build":
		err = proj.Build(flag.Arg(2))
	case "repo":
		err = proj.Repo(flag.Arg(2))
	default:
		err = &UnknownCommand{
			errors.Newf("cmd: Unknown cmd '%s'", cmd),
		}
	}

	return
}
