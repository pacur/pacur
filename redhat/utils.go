package redhat

import (
	"fmt"

	"github.com/pacur/pacur/pack"
)

func convertArch(arch string) string {
	switch arch {
	case "all":
		return "noarch"
	case "arm64":
		return "aarch64"
	case "amd64":
		return "x86_64"
	default:
		return arch
	}
}

func ConvertSection(section string) (converted string) {
	switch section {
	case "admin":
		converted = "Applications/System"
	case "localization":
		converted = "Development/Languages"
	case "mail":
		converted = "Applications/Communications"
	case "comm":
		converted = "Applications/Communications"
	case "math":
		converted = "Applications/Productivity"
	case "database":
		converted = "Applications/Databases"
	case "misc":
		converted = "Applications/System"
	case "debug":
		converted = "Development/Debuggers"
	case "net":
		converted = "Applications/Internet"
	case "news":
		converted = "Applications/Publishing"
	case "devel":
		converted = "Development/Tools"
	case "doc":
		converted = "Documentation"
	case "editors":
		converted = "Applications/Editors"
	case "electronics":
		converted = "Applications/Engineering"
	case "embedded":
		converted = "Applications/Engineering"
	case "fonts":
		converted = "Interface/Desktops"
	case "games":
		converted = "Amusements/Games"
	case "science":
		converted = "Applications/Engineering"
	case "shells":
		converted = "System Environment/Shells"
	case "sound":
		converted = "Applications/Multimedia"
	case "graphics":
		converted = "Applications/Multimedia"
	case "text":
		converted = "Applications/Text"
	case "httpd":
		converted = "Applications/Internet"
	case "vcs":
		converted = "Development/Tools"
	case "interpreters":
		converted = "Development/Tools"
	case "video":
		converted = "Applications/Multimedia"
	case "web":
		converted = "Applications/Internet"
	case "kernel":
		converted = "System Environment/Kernel"
	case "x11":
		converted = "User Interface/X"
	case "libdevel":
		converted = "Development/Libraries"
	case "libs":
		converted = "System Environment/Libraries"
	default:
		converted = section
	}

	return
}

func formatDepend(dpn *pack.Dependency) string {
	if dpn.Comparison == "" || dpn.Version == "" {
		return dpn.Name
	}
	return fmt.Sprintf("%s %s %s", dpn.Name, dpn.Comparison, dpn.Version)
}
