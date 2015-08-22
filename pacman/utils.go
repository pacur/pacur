package pacman

func convertPacman(arch string) string {
	if arch == "all" {
		return "any"
	}
	return arch
}
