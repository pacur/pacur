package debian

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Debian struct {
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

	args := []string{
		"--assume-yes",
		"install",
	}
	args = append(args, d.Pack.MakeDepends...)

	cmd := exec.Command("apt-get", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "utils: Failed to get make depends '%s'"),
		}
		return
	}

	return
}

func (d *Debian) getSums() (err error) {
	cmd := exec.Command("find", ".", "-type", "f",
		"-exec", "md5sum", "{}", ";")
	cmd.Dir = d.Pack.PackageDir
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		err = &HashError{
			errors.Wrapf(err, "debian: Failed to get md5 hashes for '%s'",
				d.Pack.PackageDir),
		}
		return
	}

	d.sums = ""
	for _, line := range strings.Split(string(output), "\n") {
		d.sums += strings.Replace(line, "./", "", 1) + "\n"
	}

	return
}

func (d *Debian) createConfFiles() (err error) {
	if len(d.Pack.Backup) == 0 {
		return
	}

	path := filepath.Join(d.debDir, "conffiles")

	file, err := os.Create(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"debian: Failed to create debian conffiles at '%s'", path),
		}
		return
	}
	defer file.Close()

	for _, name := range d.Pack.Backup {
		_, err = file.WriteString(name + "\n")
		if err != nil {
			err = &WriteError{
				errors.Wrapf(err,
					"debian: Failed to write debian conffiles at '%s'", path),
			}
			return
		}
	}

	return
}

func (d *Debian) createControl() (err error) {
	path := filepath.Join(d.debDir, "control")

	file, err := os.Create(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"debian: Failed to create debian control at '%s'", path),
		}
		return
	}
	defer file.Close()

	data := ""

	data += fmt.Sprintf("Package: %s\n", d.Pack.PkgName)
	data += fmt.Sprintf("Version: %s-0ubuntu%s~%s\n",
		d.Pack.PkgVer, d.Pack.PkgRel, d.Release)
	data += fmt.Sprintf("Architecture: %s\n", d.Pack.Arch)
	data += fmt.Sprintf("Maintainer: %s\n", d.Pack.Maintainer)
	data += fmt.Sprintf("Installed-Size: %d\n", d.installSize)
	data += fmt.Sprintf("Depends: %s\n", strings.Join(d.Pack.Depends, ", "))
	data += fmt.Sprintf("Recommends: %s\n",
		strings.Join(d.Pack.OptDepends, ", "))
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

	_, err = file.WriteString(data)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"debian: Failed to write debian control at '%s'", path),
		}
		return
	}

	return
}

func (d *Debian) createMd5Sums() (err error) {
	if len(d.Pack.Backup) == 0 {
		return
	}

	path := filepath.Join(d.debDir, "md5sums")

	file, err := os.Create(path)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"debian: Failed to create debian md5sums at '%s'", path),
		}
		return
	}
	defer file.Close()

	_, err = file.WriteString(d.sums)
	if err != nil {
		err = &WriteError{
			errors.Wrapf(err,
				"debian: Failed to write debian md5sums at '%s'", path),
		}
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

		file, e := os.Create(path)
		if e != nil {
			err = &WriteError{
				errors.Wrapf(e,
					"debian: Failed to create debian %s at '%s'", name, path),
			}
			return
		}
		defer file.Close()
		if e != nil {
			err = &WriteError{
				errors.Wrapf(e,
					"debian: Failed to chmod debian %s at '%s'", name, path),
			}
			return
		}

		err = os.Chmod(path, 0755)

		_, err = file.WriteString(strings.Join(script, "\n"))
		if err != nil {
			err = &WriteError{
				errors.Wrapf(err,
					"debian: Failed to write debian %s at '%s'", name, path),
			}
			return
		}
	}

	return
}

func (d *Debian) dpkgDeb() (err error) {
	cmd := exec.Command("dpkg-deb", "-b", d.Pack.PackageDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "debian: Failed to build dpkg-deb '%s'",
				d.Pack.PackageDir),
		}
		return
	}

	_, dir := filepath.Split(filepath.Clean(d.Pack.Root))
	path := filepath.Join(d.Pack.Root, dir+".deb")
	newPath := filepath.Join(d.Pack.Root,
		fmt.Sprintf("%s_%s-0ubuntu%s.%s_%s.deb",
			d.Pack.PkgName, d.Pack.PkgVer, d.Pack.PkgRel,
			d.Release, d.Pack.Arch))

	os.Remove(newPath)

	err = os.Rename(path, newPath)
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "debian: Failed to rename deb '%s'",
				d.Pack.PackageDir),
		}
		return
	}

	return
}

func (d *Debian) Build() (err error) {
	err = d.getDepends()
	if err != nil {
		return
	}

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
