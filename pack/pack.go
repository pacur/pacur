package pack

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/resolver"
	"strings"
)

type Dependency struct {
	Name        string
	Restriction *string
}

type Pack struct {
	priorities  map[string]int
	Targets     []string
	Distro      string
	Release     string
	FullRelease string
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
	Depends     []Dependency
	OptDepends  []Dependency
	MakeDepends []Dependency
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
	Variables   map[string]string
	RpmOpts     []string
}

func (p *Pack) Init() {
	p.priorities = map[string]int{}

	p.FullRelease = p.Distro
	if p.Release != "" {
		p.FullRelease += "-" + p.Release
	}
}

func (p *Pack) parseDirective(input string) (key string, pry int, err error) {
	split := strings.Split(input, ":")
	key = split[0]

	if len(split) == 1 {
		pry = 0
		return
	} else if len(split) != 2 {
		err = &ParseError{
			errors.Newf("pack: Invalid use of ':' directive in '%s'", input),
		}
		return
	} else {
		pry = -1
	}

	if p.Distro == "" {
		return
	}

	if key == "pkgver" || key == "pkgrel" {
		err = &ParseError{
			errors.Newf("pack: Cannot use directive for '%s'", key),
		}
		return
	}

	dirc := split[1]

	if constants.ReleasesSet.Contains(dirc) {
		if dirc == p.FullRelease {
			pry = 3
		}
		return
	}

	if constants.DistrosSet.Contains(dirc) {
		if dirc == p.Distro {
			pry = 2
		}
		return
	}

	if constants.PackagersSet.Contains(dirc) {
		if dirc == constants.DistroPackager[p.Distro] {
			pry = 1
		}
		return
	}

	err = &ParseError{
		errors.Newf("pack: Unknown directive '%s'", dirc),
	}
	return
}

func ParseDependency(dependency string) Dependency {
	for i, c := range dependency {
		if c == '=' || c == '<' || c == '>' {
			var restriction = dependency[i:]
			return Dependency {
				Name: dependency[:i],
				Restriction: &restriction,
			}
		}
	}
	return Dependency {
		Name: dependency,
		Restriction: nil,
	}
}

func ParseDependencies(dependencies []string) []Dependency {
	parsed := make([]Dependency, len(dependencies))
	for i, d := range dependencies {
		parsed[i] = ParseDependency(d)
	}
	return parsed
}

func (p *Pack) Resolve() (err error) {
	reslv := resolver.New()

	var dependsRaw     []string
	var optDependsRaw  []string
	var makeDependsRaw []string

	reslv.AddList("targets", p.Targets)
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
	reslv.AddList("depends", dependsRaw)
	reslv.AddList("optdepends", optDependsRaw)
	reslv.AddList("makedepends", makeDependsRaw)
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
	reslv.AddList("rpmopts", p.RpmOpts)

	if p.Variables != nil {
		for key, val := range p.Variables {
			reslv.Add(key, &val)
		}
	}

	err = reslv.Resolve()
	if err != nil {
		return
	}

	return
}

func (p *Pack) AddItem(key string, data interface{}, n int, line string) (
	err error) {

	key, priority, err := p.parseDirective(key)
	if err != nil {
		return
	}

	if priority == -1 {
		return
	}

	if priority < p.priorities[key] {
		return
	}
	p.priorities[key] = priority

	switch key {
	case "targets":
		p.Targets = data.([]string)
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
		p.Depends = ParseDependencies(data.([]string))
	case "optdepends":
		p.OptDepends = ParseDependencies(data.([]string))
	case "makedepends":
		p.MakeDepends = ParseDependencies(data.([]string))
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
	case "rpmopts":
		p.RpmOpts = data.([]string)
	default:
		if p.Variables == nil {
			p.Variables = map[string]string{}
		}
		p.Variables[key] = data.(string)
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
