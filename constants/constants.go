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
	DistroPack = map[string]string{
		"archlinux": "pacman",
		"fedora":    "redhat",
		"centos":    "redhat",
		"debian":    "debian",
	}
)
