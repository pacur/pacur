package redhat

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/dropbox/godropbox/container/set"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/utils"
	"io/ioutil"
	"os"
	"os/exec"
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

func (r *Redhat) getDepends() (err error) {
	if len(r.Pack.MakeDepends) == 0 {
		return
	}

	args := []string{
		"-y",
		"install",
	}
	args = append(args, r.Pack.MakeDepends...)

	cmd := exec.Command("yum", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "redhat: Failed to get make depends '%s'"),
		}
		return
	}

	return
}

func (r *Redhat) getFiles() (files []string, err error) {
	filesSet := set.NewSet()
	backups := set.NewSet()

	for _, path := range r.Pack.Backup {
		backups.Add(path)
	}

	cmd := exec.Command("find", ".", "-type", "f")
	cmd.Dir = r.Pack.PackageDir
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "redhat: Failed to get file list for '%s'",
				r.Pack.PackageDir),
		}
		return
	}

	for _, path := range strings.Split(string(output), "\n") {
		if len(path) < 1 {
			continue
		}
		path = path[1:]

		if strings.HasSuffix(path, ".py") || strings.HasSuffix(path, ".pyc") ||
			strings.HasSuffix(path, ".pyo") {
			path = path[:strings.LastIndex(path, ".")] + ".*"
		}

		path = `"` + path + `"`

		if filesSet.Contains(path) {
			continue
		}

		if backups.Contains(path) {
			path = "%config " + path
		}

		files = append(files, path)
		filesSet.Add(path)
	}

	cmd = exec.Command("find", ".", "-type", "d", "-empty")
	cmd.Dir = r.Pack.PackageDir
	cmd.Stderr = os.Stderr

	output, err = cmd.Output()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "redhat: Failed to get dir list for '%s'",
				r.Pack.PackageDir),
		}
		return
	}

	for _, path := range strings.Split(string(output), "\n") {
		if len(path) < 1 {
			continue
		}
		path = "%dir " + path[1:]
		path = `"` + path + `"`
		files = append(files, path)
	}

	return
}

func (r *Redhat) createSpec() (err error) {
	path := filepath.Join(r.specsDir, r.Pack.PkgName+".spec")

	file, err := os.Create(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"redhat: Failed to create redhat spec at '%s'", path),
		}
		return
	}
	defer file.Close()

	data := ""

	data += fmt.Sprintf("Name: %s\n", r.Pack.PkgName)
	data += fmt.Sprintf("Summary: %s\n", r.Pack.PkgDesc)
	data += fmt.Sprintf("Version: %s\n", r.Pack.PkgVer)
	data += fmt.Sprintf("Release: %s", r.Pack.PkgName) + "%{?dist}\n"
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
	data += fmt.Sprintf("mv -f %s/.[!.]* $RPM_BUILD_ROOT || true\n",
		r.Pack.PackageDir)
	data += fmt.Sprintf("mv -f %s/* $RPM_BUILD_ROOT || true\n",
		r.Pack.PackageDir)
	data += "\n"

	files, err := r.getFiles()
	if err != nil {
		return
	}

	data += "%files\n"
	for _, line := range files {
		data += line + "\n"
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

	_, err = file.WriteString(data)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"redhat: Failed to write redhat spec at '%s'", path),
		}
		return
	}

	fmt.Println(data)

	return
}

func (r *Redhat) rpmBuild() (err error) {
	cmd := exec.Command("rpmbuild", "--define", "_topdir "+r.redhatDir,
		"-ba", r.Pack.PkgName+".spec")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = r.specsDir

	err = cmd.Run()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "redhat: Failed to run rpmbuild '%s'",
				r.redhatDir),
		}
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

func (r *Redhat) Build() (err error) {
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
	defer os.RemoveAll(r.redhatDir)

	err = r.createSpec()
	if err != nil {
		return
	}

	err = r.rpmBuild()
	if err != nil {
		return
	}

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
			r.Pack.Home)
		if err != nil {
			return
		}
	}

	err = utils.CopyFiles(r.srpmsDir, r.Pack.Home)
	if err != nil {
		return
	}

	return
}
