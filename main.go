package main

import (
	"flag"
	"github.com/pacur/pacur/builder"
	"github.com/pacur/pacur/debian"
	"github.com/pacur/pacur/parse"
)

func main() {
	flag.Parse()

	distro := ""
	release := ""

	arg := flag.Arg(0)
	switch arg {
	case "ubuntu-precise":
		distro = "ubuntu"
		release = "precise"
	case "ubuntu-trusty":
		distro = "ubuntu"
		release = "trusty"
	case "ubuntu-vivid":
		distro = "ubuntu"
		release = "vivid"
	case "ubuntu-wily":
		distro = "ubuntu"
		release = "wily"
	default:
		panic("main: Unknown build distro or release " + arg)
	}

	pac, err := parse.File("/pacur/PKGBUILD")
	if err != nil {
		panic(err)
	}

	err = pac.Compile()
	if err != nil {
		panic(err)
	}

	if distro != "ubuntu" {
		panic("main: Unknown distro")
	}

	deb := debian.Debian{
		Pack:    pac,
		Release: release,
	}

	err = deb.Prep()
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

	err = deb.Build()
	if err != nil {
		panic(err)
	}

	//	switch distro {
	//	case "ubuntu":
	//		deb := debian.Debian{
	//			Pack:    pac,
	//			Release: release,
	//		}
	//		err = deb.Build()
	//		if err != nil {
	//			panic(err)
	//		}
	//	default:
	//		panic("main: Unknown distro")
	//	}
}
