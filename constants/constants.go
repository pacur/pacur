package constants

import (
	"strings"

	"github.com/dropbox/godropbox/container/set"
)

const (
	DockerOrg = "pacur/"
)

var (
	Releases = [...]string{
		"archlinux",
		"amazonlinux-2016.09",
		"amazonlinux-2017.03",
		"fedora-24",
		"fedora-25",
		//		"fedora-26",
		"centos-7",
		"debian-wheezy",
		"debian-jessie",
		"debian-stretch",
		"ubuntu-precise",
		"ubuntu-trusty",
		"ubuntu-xenial",
		"ubuntu-yakkety",
		"ubuntu-zesty",
	}
	ReleasesMatch = map[string]string{
		"archlinux":           "",
		"amazonlinux-2016.09": ".amzn.2016.09.",
		"amazonlinux-2017.03": ".amzn.2017.03.",
		"fedora-24":           ".fc24.",
		"fedora-25":           ".fc25.",
		"fedora-26":           ".fc26.",
		"centos-6":            ".el6.centos.",
		"centos-7":            ".el7.centos.",
		"debian-wheezy":       ".wheezy_",
		"debian-jessie":       ".jessie_",
		"debian-stretch":      ".stretch_",
		"ubuntu-precise":      ".precise_",
		"ubuntu-trusty":       ".trusty_",
		"ubuntu-xenial":       ".xenial_",
		"ubuntu-yakkety":      ".yakkety_",
		"ubuntu-zesty":        ".zesty_",
	}
	DistroPack = map[string]string{
		"archlinux":   "pacman",
		"amazonlinux": "redhat",
		"fedora":      "redhat",
		"centos":      "redhat",
		"debian":      "debian",
		"ubuntu":      "debian",
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
