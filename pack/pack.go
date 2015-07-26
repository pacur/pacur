package pack

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/resolver"
)

type Pack struct {
	Root        string
	Home        string
	SourceDir   string
	PackageDir  string
	PkgName     string
	PkgVer      string
	PkgRel      string
	PkgDesc     string
	PkgDescLong []string
	Maintainer  string
	Arch        string
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

func (p *Pack) Resolve() (err error) {
	reslv := resolver.New()

	reslv.Add("root", &p.Root)
	reslv.Add("srcdir", &p.SourceDir)
	reslv.Add("pkgdir", &p.PackageDir)
	reslv.Add("pkgname", &p.PkgName)
	reslv.Add("pkgver", &p.PkgVer)
	reslv.Add("pkgrel", &p.PkgRel)
	reslv.Add("pkgdesc", &p.PkgDesc)
	reslv.AddList("pkgdesclong", p.PkgDescLong)
	reslv.Add("maintainer", &p.Maintainer)
	reslv.Add("arch", &p.Arch)
	reslv.AddList("license", p.License)
	reslv.Add("section", &p.Section)
	reslv.Add("priority", &p.Priority)
	reslv.Add("url", &p.Url)
	reslv.AddList("depends", p.Depends)
	reslv.AddList("optdepends", p.OptDepends)
	reslv.AddList("makedepends", p.MakeDepends)
	reslv.AddList("provides", p.Provides)
	reslv.AddList("conflicts", p.Conflicts)
	reslv.AddList("sources", p.Sources)
	reslv.AddList("hashsums", p.HashSums)
	reslv.AddList("backup", p.Backup)
	reslv.AddList("build", p.Build)
	reslv.AddList("package", p.Package)
	reslv.AddList("preinst", p.PreInst)
	reslv.AddList("postinst", p.PostInst)
	reslv.AddList("prerm", p.PreRm)
	reslv.AddList("postrm", p.PostRm)

	err = reslv.Resolve()
	if err != nil {
		return
	}

	return
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
		p.Arch = data.(string)
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

func (p *Pack) Validate() (err error) {
	if len(p.Sources) == len(p.HashSums) {
	} else if len(p.Sources) > len(p.HashSums) {
		err = &ValidationError{
			errors.New("pack: Missing hash sum for source"),
		}
		return
	} else {
		err = &ValidationError{
			errors.New("pack: Too many hash sums for sources"),
		}
		return
	}

	return
}

func (p *Pack) Compile() (err error) {
	err = p.Validate()
	if err != nil {
		return
	}

	err = p.Resolve()
	if err != nil {
		return
	}

	return
}
