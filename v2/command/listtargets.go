package command

import (
	"fmt"
	"github.com/aanno/pacur/v2/constants"
)

func ListTargets() (_ error) {
	for _, release := range constants.Releases {
		fmt.Println(release)
	}

	return
}
