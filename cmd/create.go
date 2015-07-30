package cmd

import (
	"flag"
	"github.com/pacur/pacur/mirror"
	"strings"
)

func Create() (err error) {
	split := strings.Split(flag.Arg(1), "-")
	distro := split[0]
	release := ""
	if len(split) > 1 {
		release = split[1]
	}

	mirr := &mirror.Mirror{
		Root:    "/pacur",
		Distro:  distro,
		Release: release,
	}

	err = mirr.Create()
	if err != nil {
		return
	}

	return
}
