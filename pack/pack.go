package pack

import (
	"github.com/dropbox/godropbox/errors"
)

type Pack struct {
	Root        string
	PkgName     string
	PkgVer      string
	PkgRel      string
	PkgDesc     string
	PkgDescLong []string
	Maintainer  string
	Arch        []string
	License     []string
	Section     string
	Priority    string
	Url         string
	Depends     []string
	OptDepends  []string
	MakeDepends []string
	Provides    []string
	Conflicts   []string
	Sources     []string
	HashSums    []string
	Backup      []string
	Build       []string
	Package     []string
	PreInst     []string
	PostInst    []string
	PreRm       []string
	PostRm      []string
}

func (p *Pack) AddItem(key string, data interface{}, n int, line string) (
	err error) {

	switch key {
	case "pkgname":
		p.PkgName = data.(string)
	case "pkgver":
		p.PkgVer = data.(string)
	case "pkgrel":
		p.PkgRel = data.(string)
	case "pkgdesc":
		p.PkgDesc = data.(string)
	case "pkgdesclong":
		p.PkgDescLong = data.([]string)
	case "maintainer":
		p.Maintainer = data.(string)
	case "arch":
		p.Arch = data.([]string)
	case "license":
		p.License = data.([]string)
	case "section":
		p.Section = data.(string)
	case "priority":
		p.Priority = data.(string)
	case "url":
		p.Url = data.(string)
	case "depends":
		p.Depends = data.([]string)
	case "optdepends":
		p.OptDepends = data.([]string)
	case "makedepends":
		p.MakeDepends = data.([]string)
	case "provides":
		p.Provides = data.([]string)
	case "conflicts":
		p.Conflicts = data.([]string)
	case "sources":
		p.Sources = data.([]string)
	case "hashsums":
		p.HashSums = data.([]string)
	case "backup":
		p.Backup = data.([]string)
	case "build":
		p.Build = data.([]string)
	case "package":
		p.Package = data.([]string)
	case "preinst":
		p.PreInst = data.([]string)
	case "postinst":
		p.PostInst = data.([]string)
	case "prerm":
		p.PreRm = data.([]string)
	case "postrm":
		p.PostRm = data.([]string)
	default:
		err = &ParseError{
			errors.Newf("pack: Unknown option '%s' (%d: %s)",
				key, n, line),
		}
	}

	return
}
