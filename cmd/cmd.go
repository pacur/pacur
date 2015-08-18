package cmd

import (
	"flag"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
)

func Parse() (err error) {
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "build":
		err = Build()
	case "create":
		err = Create()
	case "project":
		err = Project()
	case "pull":
		err = utils.PullContainers()
	default:
		err = &UnknownCommand{
			errors.Newf("cmd: Unknown command '%s'", cmd),
		}
	}

	return
}
