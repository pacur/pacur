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
		"almalinux-8",
		"almalinux-9",
		"amazonlinux-1",
		"amazonlinux-2",
		"amazonlinux-2023",
		"fedora-38",
		"fedora-39",
		"centos-7",
		"centos-8",
		"debian-buster",
		"debian-bullseye",
		"debian-bookworm",
		"oraclelinux-7",
		"oraclelinux-8",
		"oraclelinux-9",
		"ubuntu-xenial",
		"ubuntu-bionic",
		"ubuntu-focal",
		"ubuntu-jammy",
		"ubuntu-mantic",
		"ubuntu-noble",
	}
	ReleasesMatch = map[string]string{
		"archlinux":        "",
		"almalinux-8":      ".el8.almalinux.",
		"almalinux-9":      ".el9.almalinux.",
		"amazonlinux-1":    ".amzn1.",
		"amazonlinux-2":    ".amzn2.",
		"amazonlinux-2023": ".amzn2023.",
		"fedora-38":        ".fc38.",
		"fedora-39":        ".fc39.",
		"centos-7":         ".el7.centos.",
		"centos-8":         ".el8.centos.",
		"debian-buster":    ".buster_",
		"debian-bullseye":  ".bullseye_",
		"debian-bookworm":  ".bookworm_",
		"oraclelinux-7":    ".el7.oraclelinux.",
		"oraclelinux-8":    ".el8.oraclelinux.",
		"oraclelinux-9":    ".el9.oraclelinux.",
		"ubuntu-xenial":    ".xenial_",
		"ubuntu-bionic":    ".bionic_",
		"ubuntu-focal":     ".focal_",
		"ubuntu-jammy":     ".jammy_",
		"ubuntu-mantic":    ".mantic_",
		"ubuntu-noble":     ".noble_",
	}
	DistroPack = map[string]string{
		"archlinux":   "pacman",
		"almalinux":   "redhat",
		"amazonlinux": "redhat",
		"fedora":      "redhat",
		"centos":      "redhat",
		"debian":      "debian",
		"oraclelinux": "redhat",
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
	CleanPrevious  = false
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
