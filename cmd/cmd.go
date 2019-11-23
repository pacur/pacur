package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/m0rf30/pacur/utils"
	"github.com/urfave/cli"
)

// Parse retrieve cli strings and
// return an error
func Parse() (err error) {
	flag.Parse()

	app := cli.NewApp()
	app.Name = "pacur"
	app.Usage = "Automated deb, rpm and pkgbuild build system'"

	app.Commands = []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Build a project",
			Action: func(c *cli.Context) error {
				err = Build()
				return err
			},
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create a project",
			Action: func(c *cli.Context) error {
				err = Create()
				return err
			},
		},
		{
			Name:    "project",
			Aliases: []string{"p"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				err = Project()
				return err
			},
		},
		{
			Name:    "docker",
			Aliases: []string{"d"},
			Usage:   "Pull the built images",
			Action: func(c *cli.Context) error {
				err = utils.PullContainers()
				return err
			},
		},
		{
			Name:    "gen-key",
			Aliases: []string{"g"},
			Usage:   "Generate a pairs of key for repo",
			Action: func(c *cli.Context) error {
				err = GenKey()
				return err
			},
		},
		{
			Name:    "list-targets",
			Aliases: []string{"l"},
			Usage:   "List a bunch of available build targets",
			Action: func(c *cli.Context) error {
				err = ListTargets()
				return err
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
