package cmd

import (
	"flag"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/project"
	"github.com/pacur/pacur/utils"
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
	proj.Init()

	cmd := flag.Arg(1)
	switch cmd {
	case "init":
		err = proj.InitProject()
	case "pull":
		err = utils.PullContainers()
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
