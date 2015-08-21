package debian

import (
	"fmt"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/utils"
	"os"
	"path/filepath"
	"strings"
)

type Debian struct {
	Distro      string
	Release     string
	Pack        *pack.Pack
	debDir      string
	installSize int
	sums        string
}

func (d *Debian) getDepends() (err error) {
	if len(d.Pack.MakeDepends) == 0 {
		return
	}

	err = utils.Exec("", "apt-get", "--assume-yes", "update")
	if err != nil {
		return
	}

	args := []string{
		"--assume-yes",
		"install",
	}
	args = append(args, d.Pack.MakeDepends...)

	err = utils.Exec("", "apt-get", args...)
	if err != nil {
		return
	}

	return
}

func (d *Debian) getSums() (err error) {
	output, err := utils.ExecOutput(d.Pack.PackageDir, "find", ".",
		"-type", "f", "-exec", "md5sum", "{}", ";")
	if err != nil {
		return
	}

	d.sums = ""
	for _, line := range strings.Split(output, "\n") {
		d.sums += strings.Replace(line, "./", "", 1) + "\n"
	}

	return
}

func (d *Debian) createConfFiles() (err error) {
	if len(d.Pack.Backup) == 0 {
		return
	}

	path := filepath.Join(d.debDir, "conffiles")

	data := ""
	for _, name := range d.Pack.Backup {
		data += name + "\n"
	}

	err = utils.CreateWrite(path, data)
	if err != nil {
		return
	}

	return
}

func (d *Debian) createControl() (err error) {
	path := filepath.Join(d.debDir, "control")

	data := ""

	data += fmt.Sprintf("Package: %s\n", d.Pack.PkgName)
	data += fmt.Sprintf("Version: %s-0%s%s~%s\n",
		d.Pack.PkgVer, d.Distro, d.Pack.PkgRel, d.Release)
	data += fmt.Sprintf("Architecture: %s\n", d.Pack.Arch)
	data += fmt.Sprintf("Maintainer: %s\n", d.Pack.Maintainer)
	data += fmt.Sprintf("Installed-Size: %d\n", d.installSize)

	if len(d.Pack.Depends) > 0 {
		data += fmt.Sprintf("Depends: %s\n",
			strings.Join(d.Pack.Depends, ", "))
	}

	if len(d.Pack.OptDepends) > 0 {
		data += fmt.Sprintf("Recommends: %s\n",
			strings.Join(d.Pack.OptDepends, ", "))
	}

	data += fmt.Sprintf("Section: %s\n", d.Pack.Section)
	data += fmt.Sprintf("Priority: %s\n", d.Pack.Priority)
	data += fmt.Sprintf("Homepage: %s\n", d.Pack.Url)
	data += fmt.Sprintf("Description: %s\n", d.Pack.PkgDesc)

	for _, line := range d.Pack.PkgDescLong {
		if line == "" {
			line = "."
		}
		data += fmt.Sprintf("  %s\n", line)
	}

	err = utils.CreateWrite(path, data)
	if err != nil {
		return
	}

	return
}

func (d *Debian) createMd5Sums() (err error) {
	path := filepath.Join(d.debDir, "md5sums")

	err = utils.CreateWrite(path, d.sums)
	if err != nil {
		return
	}

	return
}

func (d *Debian) createScripts() (err error) {
	scripts := map[string][]string{
		"preinst":  d.Pack.PreInst,
		"postinst": d.Pack.PostInst,
		"prerm":    d.Pack.PreRm,
		"postrm":   d.Pack.PostRm,
	}

	for name, script := range scripts {
		if len(script) == 0 {
			continue
		}

		path := filepath.Join(d.debDir, name)

		err = utils.CreateWrite(path, strings.Join(script, "\n"))
		if err != nil {
			return
		}

		err = utils.Chmod(path, 0755)
		if err != nil {
			return
		}
	}

	return
}

func (d *Debian) dpkgDeb() (err error) {
	err = utils.Exec("", "dpkg-deb", "-b", d.Pack.PackageDir)
	if err != nil {
		return
	}

	_, dir := filepath.Split(filepath.Clean(d.Pack.PackageDir))
	path := filepath.Join(d.Pack.Root, dir+".deb")
	newPath := filepath.Join(d.Pack.Home,
		fmt.Sprintf("%s_%s-0%s%s.%s_%s.deb",
			d.Pack.PkgName, d.Pack.PkgVer, d.Distro, d.Pack.PkgRel,
			d.Release, d.Pack.Arch))

	os.Remove(newPath)

	err = utils.CopyFile("", path, newPath, false)
	if err != nil {
		return
	}

	return
}

func (d *Debian) Prep() (err error) {
	err = d.getDepends()
	if err != nil {
		return
	}

	return
}

func (d *Debian) Build() (err error) {
	d.installSize, err = utils.GetDirSize(d.Pack.PackageDir)
	if err != nil {
		return
	}

	err = d.getSums()
	if err != nil {
		return
	}

	d.debDir = filepath.Join(d.Pack.PackageDir, "DEBIAN")
	err = utils.ExistsMakeDir(d.debDir)
	if err != nil {
		return
	}
	defer os.RemoveAll(d.debDir)

	err = d.createConfFiles()
	if err != nil {
		return
	}

	err = d.createControl()
	if err != nil {
		return
	}

	err = d.createMd5Sums()
	if err != nil {
		return
	}

	err = d.createScripts()
	if err != nil {
		return
	}

	err = d.dpkgDeb()
	if err != nil {
		return
	}

	return
}
