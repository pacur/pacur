package pacman

import (
	"fmt"
	"github.com/pacur/pacur/pack"
)

func convertPacman(arch string) string {
	switch arch {
	case "all":
		return "any"
	case "amd64":
		return "x86_64"
	default:
		return arch
	}
}

func formatDepend(dpn *pack.Dependency) string {
	if dpn.Comparison == "" || dpn.Version == "" {
		return dpn.Name
	}
	return fmt.Sprintf("%s%s%s", dpn.Name, dpn.Comparison, dpn.Version)
}
