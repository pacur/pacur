package main

import (
	"flag"
	"github.com/pacur/pacur/builder"
	"github.com/pacur/pacur/packer"
	"github.com/pacur/pacur/parse"
	"strings"
)

func main() {
	flag.Parse()

	split := strings.Split(flag.Arg(0), "-")
	distro := split[0]
	release := ""
	if len(split) > 1 {
		release = split[1]
	}

	pac, err := parse.File("/pacur")
	if err != nil {
		panic(err)
	}

	err = pac.Compile()
	if err != nil {
		panic(err)
	}

	pcker, err := packer.GetPacker(pac, distro, release)
	if err != nil {
		panic(err)
	}

	err = pcker.Prep()
	if err != nil {
		panic(err)
	}

	builder := builder.Builder{
		Pack: pac,
	}
	err = builder.Build()
	if err != nil {
		panic(err)
	}

	err = pcker.Build()
	if err != nil {
		panic(err)
	}
}
