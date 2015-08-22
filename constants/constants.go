package constants

const (
	DockerOrg = "pacur/"
)

var (
	Releases = [...]string{
		"archlinux",
		"centos-7",
		"debian-jessie",
		"debian-wheezy",
		"ubuntu-precise",
		"ubuntu-trusty",
		"ubuntu-vivid",
		"ubuntu-wily",
	}
	BaseDistro = map[string]string{
		"archlinux":      "pacman",
		"centos-7":       "redhat",
		"debian-jessie":  "debian",
		"debian-wheezy":  "debian",
		"ubuntu-precise": "debian",
		"ubuntu-trusty":  "debian",
		"ubuntu-vivid":   "debian",
		"ubuntu-wily":    "debian",
	}
)
