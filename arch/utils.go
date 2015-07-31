package arch

func convertArch(arch string) string {
	if arch == "all" {
		return "any"
	}
	return arch
}
