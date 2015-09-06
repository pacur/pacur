package constants

import (
	"github.com/dropbox/godropbox/container/set"
	"strings"
)

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
	ReleasesMatch = map[string]string{
		"archlinux":      "",
		"fedora-21":      ".fc21.",
		"fedora-22":      ".fc22.",
		"centos-6":       ".el6.centos.",
		"centos-7":       ".el7.centos.",
		"debian-jessie":  ".jessie_",
		"debian-wheezy":  ".wheezy_",
		"ubuntu-precise": ".precise_",
		"ubuntu-trusty":  ".trusty_",
		"ubuntu-vivid":   ".vivid_",
		"ubuntu-wily":    ".wily_",
	}
	DistroPack = map[string]string{
		"archlinux": "pacman",
		"fedora":    "redhat",
		"centos":    "redhat",
		"debian":    "debian",
		"ubuntu":    "debian",
	}
	Packagers = [...]string{
		"apt",
		"pacman",
		"yum",
	}

	ReleasesSet    = set.NewSet()
	Distros        = []string{}
	DistrosSet     = set.NewSet()
	DistroPackager = map[string]string{}
	PackagersSet   = set.NewSet()
)

func init() {
	for _, release := range Releases {
		ReleasesSet.Add(release)
		distro := strings.Split(release, "-")[0]
		Distros = append(Distros, distro)
		DistrosSet.Add(distro)
	}

	for _, distro := range Distros {
		packager := ""

		switch DistroPack[distro] {
		case "debian":
			packager = "apt"
		case "pacman":
			packager = "pacman"
		case "redhat":
			packager = "yum"
		default:
			panic("Failed to find packager for distro")
		}

		DistroPackager[distro] = packager
	}

	for _, packager := range Packagers {
		PackagersSet.Add(packager)
	}
}
