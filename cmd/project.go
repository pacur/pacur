package cmd

import (
	"flag"
	"os"

	"github.com/m0rf30/pacur/project"
)

func Project() (err error) {
	path, err := os.Getwd()
	if err != nil {
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
	}

	return
}
