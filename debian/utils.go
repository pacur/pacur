package debian

import (
	"fmt"
	"github.com/pacur/pacur/pack"
	"strings"
)

func formatDepend(dpn *pack.Dependency) string {
	if dpn.Comparison == "" || dpn.Version == "" {
		return dpn.Name
	}
	return fmt.Sprintf("%s (%s %s)", dpn.Name, dpn.Comparison, dpn.Version)
}

func formatDepends(dpns []*pack.Dependency) string {
	dpnsStr := []string{}

	for _, dpn := range dpns {
		dpnsStr = append(dpnsStr, formatDepend(dpn))
	}

	return strings.Join(dpnsStr, ", ")
}
