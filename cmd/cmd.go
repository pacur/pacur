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
	app.Version = "0.5"

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
			Usage:   "Create and manage a project",
			Subcommands: []*cli.Command{
				{
					Name:  "init",
					Usage: "Initialize the repo and create pacur.json",
					Action: func(c *cli.Context) error {
						Project()
						return nil
					},
				},
				{
					Name:  "build",
					Usage: "Build a project",
					Action: func(c *cli.Context) error {
						Project()
						return nil
					},
				},
				{
					Name:  "repo",
					Usage: "Generate all the assets needed for repo hosting",
					Action: func(c *cli.Context) error {
						Project()
						return nil
					},
				},
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
