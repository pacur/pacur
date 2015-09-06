package project

import (
	"github.com/dropbox/godropbox/container/set"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/parse"
	"strings"
)

func getDistro(name string) (distro, release string) {
	split := strings.Split(name, "-")
	distro = split[0]
	if len(split) > 1 {
		release = split[1]
	}

	return
}

func getTargets(path string) (targets []string, err error) {
	pac, err := parse.File("", "", "/pacur")
	if err != nil {
		return
	}

	err = pac.Compile()
	if err != nil {
		return
	}

	tarSet := set.NewSet()
	for _, target := range pac.Targets {
		tarSet.Add(target)
	}

	for _, release := range constants.Releases {
		distro := strings.Split(release, "-")[0]

		if tarSet.Contains("!"+release) || tarSet.Contains("!"+distro) {
			continue
		}

		if tarSet.Contains(release) || tarSet.Contains(distro) {
			targets = append(targets, release)
		}
	}

	return
}
