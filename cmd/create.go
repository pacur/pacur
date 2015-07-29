package cmd

import (
	"flag"
	"github.com/pacur/pacur/repo"
)

func Create() (err error) {
	rpo := &repo.Repo{
		Root: "/pacur",
	}

	err = rpo.Create(flag.Arg(1))
	if err != nil {
		return
	}

	return
}
