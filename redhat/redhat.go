package redhat

import (
	"fmt"
	"github.com/dropbox/godropbox/container/set"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/constants"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Redhat struct {
	Distro       string
	Release      string
	Pack         *pack.Pack
	redhatDir    string
	buildDir     string
	buildRootDir string
	rpmsDir      string
	sourcesDir   string
	specsDir     string
	srpmsDir     string
}

func (r *Redhat) getRpmPath() (path string, err error) {
	archs, err := ioutil.ReadDir(r.rpmsDir)
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "redhat: Failed to find rpms arch from '%s'",
				r.rpmsDir),
		}
		return
	}

	if len(archs) < 1 {
		err = &BuildError{
			errors.Newf("redhat: Failed to find rpm arch from '%s'",
				r.rpmsDir),
		}
		return
	}
	archPath := filepath.Join(r.rpmsDir, archs[0].Name())

	rpms, err := ioutil.ReadDir(archPath)
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "redhat: Failed to find rpms from '%s'",
				r.rpmsDir),
		}
		return
	}

	if len(rpms) < 1 {
		err = &BuildError{
			errors.Newf("redhat: Failed to find rpm from '%s'"),
		}
		return
	}
	path = filepath.Join(archPath, rpms[0].Name())

	return
}

func (r *Redhat) getDepends() (err error) {
	if len(r.Pack.MakeDepends) == 0 {
		return
	}

	args := []string{
		"-y",
		"install",
	}
	args = append(args, r.Pack.MakeDepends...)

	err = utils.Exec("", "yum", args...)
	if err != nil {
		return
	}

	return
}

func (r *Redhat) getFiles() (files []string, err error) {
	backup := set.NewSet()
	paths := set.NewSet()

	for _, path := range r.Pack.Backup {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		backup.Add(path)
	}

	rpmPath, err := r.getRpmPath()
	if err != nil {
		return
	}

	output, err := utils.ExecOutput(r.Pack.PackageDir, "rpm", "-qlp", rpmPath)
	if err != nil {
		return
	}

	for _, path := range strings.Split(output, "\n") {
		if len(path) < 1 {
			continue
		}

		paths.Remove(filepath.Dir(path))
		paths.Add(path)
	}

	for pathInf := range paths.Iter() {
		path := pathInf.(string)

		if backup.Contains(path) {
			path = `%config "` + path + `"`
		} else {
			path = `"` + path + `"`
		}

		files = append(files, path)
	}

	return
}

func (r *Redhat) createSpec(files []string) (err error) {
	path := filepath.Join(r.specsDir, r.Pack.PkgName+".spec")

	data := ""
	data += fmt.Sprintf("Name: %s\n", r.Pack.PkgName)
	data += fmt.Sprintf("Summary: %s\n", r.Pack.PkgDesc)
	data += fmt.Sprintf("Version: %s\n", r.Pack.PkgVer)
	data += fmt.Sprintf("Release: %s", r.Pack.PkgRel) + "%{?dist}\n"
	data += fmt.Sprintf("Group: %s\n", ConvertSection(r.Pack.Section))
	data += fmt.Sprintf("URL: %s\n", r.Pack.Url)
	data += fmt.Sprintf("License: %s\n", r.Pack.License)
	data += fmt.Sprintf("Packager: %s\n", r.Pack.Maintainer)

	for _, pkg := range r.Pack.Provides {
		data += fmt.Sprintf("Provides: %s\n", pkg)
	}

	for _, pkg := range r.Pack.Conflicts {
		data += fmt.Sprintf("Conflicts: %s\n", pkg)
	}

	for _, pkg := range r.Pack.Depends {
		data += fmt.Sprintf("Requires: %s\n", pkg)
	}

	for _, pkg := range r.Pack.MakeDepends {
		data += fmt.Sprintf("BuildRequires: %s\n", pkg)
	}

	data += "\n"

	if len(r.Pack.PkgDescLong) > 0 {
		data += "%description\n"
		for _, line := range r.Pack.PkgDescLong {
			data += line + "\n"
		}
		data += "\n"
	}

	data += "%install\n"
	data += fmt.Sprintf("rsync -a -A -X %s/ $RPM_BUILD_ROOT/\n",
		r.Pack.PackageDir)
	data += "\n"

	data += "%files\n"
	if len(files) == 0 {
		data += "/\n"
	} else {
		for _, line := range files {
			data += line + "\n"
		}
	}
	data += "\n"

	if len(r.Pack.PreInst) > 0 {
		data += "%pre\n"
		for _, line := range r.Pack.PreInst {
			data += line + "\n"
		}
		data += "\n"
	}

	if len(r.Pack.PostInst) > 0 {
		data += "%post\n"
		for _, line := range r.Pack.PostInst {
			data += line + "\n"
		}
		data += "\n"
	}

	if len(r.Pack.PreRm) > 0 {
		data += "%preun\n"
		for _, line := range r.Pack.PreRm {
			data += line + "\n"
		}
		data += "\n"
	}

	if len(r.Pack.PostRm) > 0 {
		data += "%postun\n"
		for _, line := range r.Pack.PostRm {
			data += line + "\n"
		}
	}

	err = utils.CreateWrite(path, data)
	if err != nil {
		return
	}

	fmt.Println(data)

	return
}

func (r *Redhat) rpmBuild() (err error) {
	err = utils.Exec(r.specsDir, "rpmbuild", "--define",
		"_topdir "+r.redhatDir, "-ba", r.Pack.PkgName+".spec")
	if err != nil {
		return
	}

	return
}

func (r *Redhat) Prep() (err error) {
	err = r.getDepends()
	if err != nil {
		return
	}

	return
}

func (r *Redhat) makeDirs() (err error) {
	r.redhatDir = filepath.Join(r.Pack.Root, "redhat")
	r.buildDir = filepath.Join(r.redhatDir, "BUILD")
	r.buildRootDir = filepath.Join(r.redhatDir, "BUILDROOT")
	r.rpmsDir = filepath.Join(r.redhatDir, "RPMS")
	r.sourcesDir = filepath.Join(r.redhatDir, "SOURCES")
	r.specsDir = filepath.Join(r.redhatDir, "SPECS")
	r.srpmsDir = filepath.Join(r.redhatDir, "SRPMS")

	for _, path := range []string{
		r.redhatDir,
		r.buildDir,
		r.buildRootDir,
		r.rpmsDir,
		r.sourcesDir,
		r.specsDir,
		r.srpmsDir,
	} {
		err = utils.ExistsMakeDir(path)
		if err != nil {
			return
		}
	}

	return
}

func (r *Redhat) clean() (err error) {
	pkgPaths, err := utils.FindExt(r.Pack.Home, ".rpm")
	if err != nil {
		return
	}

	match, ok := constants.ReleasesMatch[r.Pack.FullRelease]
	if !ok {
		err = &BuildError{
			errors.Newf("redhat: Failed to find match for '%s'",
				r.Pack.FullRelease),
		}
		return
	}

	for _, pkgPath := range pkgPaths {
		if strings.Contains(filepath.Base(pkgPath), match) {
			_ = utils.Remove(pkgPath)
		}
	}

	return
}

func (r *Redhat) copy() (err error) {
	archs, err := ioutil.ReadDir(r.rpmsDir)
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "redhat: Failed to find rpms from '%s'",
				r.rpmsDir),
		}
		return
	}

	for _, arch := range archs {
		err = utils.CopyFiles(filepath.Join(r.rpmsDir, arch.Name()),
			r.Pack.Home, false)
		if err != nil {
			return
		}
	}

	return
}

func (r *Redhat) remDirs() {
	os.RemoveAll(r.redhatDir)
}

func (r *Redhat) Build() (err error) {
	err = r.makeDirs()
	if err != nil {
		return
	}
	defer r.remDirs()

	err = r.createSpec([]string{})
	if err != nil {
		return
	}

	err = r.rpmBuild()
	if err != nil {
		return
	}

	files, err := r.getFiles()
	if err != nil {
		return
	}

	if len(files) == 0 {
		err = &BuildError{
			errors.Newf("redhat: Failed to find rpms files '%s'",
				r.rpmsDir),
		}
		return
	}

	r.remDirs()
	err = r.makeDirs()
	if err != nil {
		return
	}

	err = r.createSpec(files)
	if err != nil {
		return
	}

	err = r.rpmBuild()
	if err != nil {
		return
	}

	err = r.clean()
	if err != nil {
		return
	}

	err = r.copy()
	if err != nil {
		return
	}

	return
}
