package constants

import (
	"strings"

	"github.com/pritunl/tools/set"
)

const (
	DockerOrg = "pacur/"
)

var (
	Releases = [...]string{
		"archlinux",
		"almalinux-8",
		"almalinux-9",
		"almalinux-10",
		"amazonlinux-2",
		"amazonlinux-2023",
		"fedora-40",
		"fedora-41",
		"fedora-42",
		"debian-buster",
		"debian-bullseye",
		"debian-bookworm",
		"debian-trixie",
		"oraclelinux-7",
		"oraclelinux-8",
		"oraclelinux-9",
		"oraclelinux-10",
		"rockylinux-8",
		"rockylinux-9",
		"rockylinux-10",
		"ubuntu-bionic",
		"ubuntu-focal",
		"ubuntu-jammy",
		"ubuntu-noble",
		"ubuntu-oracular",
		"ubuntu-plucky",
	}
	ReleasesMatch = map[string]string{
		"archlinux":        "",
		"almalinux-8":      ".el8.almalinux.",
		"almalinux-9":      ".el9.almalinux.",
		"almalinux-10":     ".el10.almalinux.",
		"amazonlinux-2":    ".amzn2.",
		"amazonlinux-2023": ".amzn2023.",
		"fedora-40":        ".fc40.",
		"fedora-41":        ".fc41.",
		"fedora-42":        ".fc42.",
		"debian-buster":    ".buster_",
		"debian-bullseye":  ".bullseye_",
		"debian-bookworm":  ".bookworm_",
		"debian-trixie":    ".trixie_",
		"oraclelinux-7":    ".el7.oraclelinux.",
		"oraclelinux-8":    ".el8.oraclelinux.",
		"oraclelinux-9":    ".el9.oraclelinux.",
		"oraclelinux-10":   ".el10.oraclelinux.",
		"rockylinux-8":     ".el8.rockylinux.",
		"rockylinux-9":     ".el9.rockylinux.",
		"rockylinux-10":    ".el10.rockylinux.",
		"ubuntu-bionic":    ".bionic_",
		"ubuntu-focal":     ".focal_",
		"ubuntu-jammy":     ".jammy_",
		"ubuntu-noble":     ".noble_",
		"ubuntu-oracular":  ".oracular_",
		"ubuntu-plucky":    ".plucky_",
	}
	DistroPack = map[string]string{
		"archlinux":   "pacman",
		"almalinux":   "redhat",
		"amazonlinux": "redhat",
		"fedora":      "redhat",
		"debian":      "debian",
		"oraclelinux": "redhat",
		"rockylinux":  "redhat",
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
