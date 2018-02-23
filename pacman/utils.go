package pacman

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
