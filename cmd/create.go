package cmd

import (
	"flag"
	"github.com/pacur/pacur/mirror"
)

func Create() (err error) {
	mirr := &mirror.Mirror{
		Root: "/pacur",
	}

	err = mirr.Create(flag.Arg(1))
	if err != nil {
		return
	}

	return
}
