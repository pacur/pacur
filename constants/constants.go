package constants

const (
	DockerOrg = "pacur/"
)

var (
	Releases = [...]string{
		"archlinux",
		"fedora-21",
		"fedora-22",
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
		"ubuntu":    "debian",
	}
)
