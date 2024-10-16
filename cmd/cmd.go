package cmd

import (
	"flag"

	"github.com/pacur/pacur/utils"
	"github.com/pritunl/tools/errors"
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
	case "genkey":
		err = GenKey()
	case "list-targets":
		err = ListTargets()
	default:
		err = &UnknownCommand{
			errors.Newf("cmd: Unknown command '%s'", cmd),
		}
	}

	return
}
