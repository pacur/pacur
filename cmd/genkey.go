package cmd

import (
	"flag"
	"strings"

	"github.com/pacur/pacur/signing"
)

func GenKey() (err error) {
	args := flag.Args()[1:]
	n := len(args)

	name := strings.Join(args[:n-1], " ")
	email := args[n-1]

	if name == "" || email == "" {
		return
	}

	gen := &signing.GenKey{
		Root:  "/pacur",
		Name:  name,
		Email: email,
	}

	err = gen.Generate()
	if err != nil {
		return
	}

	err = gen.Export()
	if err != nil {
		return
	}

	return
}
