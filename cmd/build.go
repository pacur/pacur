package cmd

import (
	"flag"
	"github.com/pacur/pacur/builder"
	"github.com/pacur/pacur/packer"
	"github.com/pacur/pacur/parse"
	"strings"
)

func Build() (err error) {
	split := strings.Split(flag.Arg(1), "-")
	distro := split[0]
	release := ""
	if len(split) > 1 {
		release = split[1]
	}

	pac, err := parse.File(distro, release, "/pacur")
	if err != nil {
		return
	}

	err = pac.Compile()
	if err != nil {
		return
	}

	pcker, err := packer.GetPacker(pac, distro, release)
	if err != nil {
		return
	}

	err = pcker.Prep()
	if err != nil {
		return
	}

	buildr := builder.Builder{
		Pack: pac,
	}
	err = buildr.Build()
	if err != nil {
		return
	}

	err = pcker.Build()
	if err != nil {
		return
	}

	return
}
