package arch

import (
	"fmt"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/utils"
	"os"
	"path/filepath"
	"strconv"
)

type Arch struct {
	Distro  string
	Release string
	Pack    *pack.Pack
	archDir string
}

func (a *Arch) getDepends() (err error) {
	if len(a.Pack.MakeDepends) == 0 {
		return
	}

	args := []string{
		"-S",
		"--noconfirm",
	}
	args = append(args, a.Pack.MakeDepends...)

	err = utils.Exec("", "pacman", args...)
	if err != nil {
		return
	}

	return
}

func (a *Arch) createInstall() (exists bool, err error) {
	data := ""

	if len(a.Pack.PreInst) > 0 {
		data += "pre_install() {\n"
		for _, line := range a.Pack.PreInst {
			data += fmt.Sprintf("    %s\n", line)
		}
		data += "}\n"
	}

	if len(a.Pack.PostInst) > 0 {
		data += "post_install() {\n"
		for _, line := range a.Pack.PostInst {
			data += fmt.Sprintf("    %s\n", line)
		}
		data += "}\n"
	}

	if len(a.Pack.PreRm) > 0 {
		data += "pre_remove() {\n"
		for _, line := range a.Pack.PreRm {
			data += fmt.Sprintf("    %s\n", line)
		}
		data += "}\n"
	}

	if len(a.Pack.PostRm) > 0 {
		data += "post_remove() {\n"
		for _, line := range a.Pack.PostRm {
			data += fmt.Sprintf("    %s\n", line)
		}
		data += "}\n"
	}

	exists = len(data) > 0
	if exists {
		path := filepath.Join(a.archDir, a.Pack.PkgName+".install")
		err = utils.CreateWrite(path, data)
		if err != nil {
			return
		}
	}

	return
}

func (a *Arch) createMake() (err error) {
	path := filepath.Join(a.archDir, "PKGBUILD")

	installExists, err := a.createInstall()
	if err != nil {
		return
	}

	data := ""
	data += fmt.Sprintf("# Maintainer: %s\n\n", a.Pack.Maintainer)
	data += fmt.Sprintf("pkgname=%s\n", strconv.Quote(a.Pack.PkgName))
	data += fmt.Sprintf("pkgver=%s\n", strconv.Quote(a.Pack.PkgVer))
	data += fmt.Sprintf("pkgrel=%s\n", strconv.Quote(a.Pack.PkgRel))
	data += fmt.Sprintf("pkgdesc=%s\n", strconv.Quote(a.Pack.PkgDesc))
	data += fmt.Sprintf("arch=(%s)\n",
		strconv.Quote(convertArch(a.Pack.Arch)))

	data += "license=(\n"
	for _, item := range a.Pack.License {
		data += fmt.Sprintf("    %s\n", strconv.Quote(item))
	}
	data += ")\n"

	data += fmt.Sprintf("url=%s\n", strconv.Quote(a.Pack.Url))

	if len(a.Pack.Depends) > 0 {
		data += "depends=(\n"
		for _, item := range a.Pack.Depends {
			data += fmt.Sprintf("    %s\n", strconv.Quote(item))
		}
		data += ")\n"
	}

	if len(a.Pack.OptDepends) > 0 {
		data += "optdepends=(\n"
		for _, item := range a.Pack.OptDepends {
			data += fmt.Sprintf("    %s\n", strconv.Quote(item))
		}
		data += ")\n"
	}

	if len(a.Pack.Provides) > 0 {
		data += "provides=(\n"
		for _, item := range a.Pack.Provides {
			data += fmt.Sprintf("    %s\n", strconv.Quote(item))
		}
		data += ")\n"
	}

	if len(a.Pack.Conflicts) > 0 {
		data += "conflicts=(\n"
		for _, item := range a.Pack.Conflicts {
			data += fmt.Sprintf("    %s\n", strconv.Quote(item))
		}
		data += ")\n"
	}

	if installExists {
		data += fmt.Sprintf("install=%s\n",
			strconv.Quote(a.Pack.PkgName+".install"))
	}

	data += `options=("emptydirs")\n`

	if len(a.Pack.Backup) > 0 {
		data += "backup=(\n"
		for _, item := range a.Pack.Backup {
			data += fmt.Sprintf("    %s\n", strconv.Quote(item))
		}
		data += ")\n"
	}

	data += "package() {\n"
	data += fmt.Sprintf("    rsync -a -A -X %s/ ${pkgdir}/\n",
		a.Pack.PackageDir)
	data += "}\n"

	err = utils.CreateWrite(path, data)
	if err != nil {
		return
	}

	fmt.Println(data)

	return
}

func (a *Arch) archBuild() (err error) {
	err = utils.Chmod(a.archDir, 0777)
	if err != nil {
		return
	}

	err = utils.Exec(a.archDir, "sudo", "-u", "nobody", "makepkg")
	if err != nil {
		return
	}

	return
}

func (a *Arch) Prep() (err error) {
	err = a.getDepends()
	if err != nil {
		return
	}

	return
}

func (a *Arch) makeDirs() (err error) {
	a.archDir = filepath.Join(a.Pack.Root, "arch")

	err = utils.ExistsMakeDir(a.archDir)
	if err != nil {
		return
	}

	return
}

func (a *Arch) remDirs() {
	os.RemoveAll(a.archDir)
}

func (a *Arch) Build() (err error) {
	err = a.makeDirs()
	if err != nil {
		return
	}
	defer a.remDirs()

	err = a.createMake()
	if err != nil {
		return
	}

	err = a.archBuild()
	if err != nil {
		return
	}

	pkgs, err := utils.FindExt(a.archDir, ".pkg.tar.xz")
	if err != nil {
		return
	}

	for _, pkg := range pkgs {
		err = utils.CopyFile("", pkg, a.Pack.Home, false)
		if err != nil {
			return
		}
	}

	return
}
