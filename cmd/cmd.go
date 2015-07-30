package cmd

import (
	"flag"
	"github.com/dropbox/godropbox/errors"
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
	default:
		err = &UnknownCommand{
			errors.Newf("cmd: Unknown command '%s'", cmd),
		}
	}

	return
}
