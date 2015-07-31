package project

import (
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
