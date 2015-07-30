package cmd

import (
	"flag"
	"github.com/pacur/pacur/mirror"
	"strings"
)

func Create() (err error) {
	mirr := &mirror.Mirror{
		Root: "/pacur",
	}

	split := strings.Split(flag.Arg(1), "-")
	distro := split[0]
	release := ""
	if len(split) > 1 {
		release = split[1]
	}

	err = mirr.Create(distro, release)
	if err != nil {
		return
	}

	return
}
