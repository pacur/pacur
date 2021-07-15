package command

import (
	"flag"
	"github.com/aanno/pacur/v2/builder"
	"github.com/aanno/pacur/v2/packer"
	"github.com/aanno/pacur/v2/parse"
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
