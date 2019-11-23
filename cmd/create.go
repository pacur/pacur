package cmd

import (
	"flag"
	"github.com/m0rf30/pacur/mirror"
	"strings"
)

func Create() (err error) {
	split := strings.Split(flag.Arg(1), "-")
	distro := split[0]
	release := ""
	if len(split) > 1 {
		release = split[1]
	}

	name := flag.Arg(2)
	if name == "" {
		name = "pacur"
	}

	mirr := &mirror.Mirror{
		Name:    name,
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
