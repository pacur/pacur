package builder

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/source"
	"github.com/pacur/pacur/utils"
	"os"
	"os/exec"
	"path/filepath"
)

type Builder struct {
	id   string
	Pack *pack.Pack
}

func (b *Builder) initDirs() (err error) {
	err = utils.ExistsMakeDir(b.Pack.SourceDir)
	if err != nil {
		return
	}

	err = utils.ExistsMakeDir(b.Pack.PackageDir)
	if err != nil {
		return
	}

	return
}

func (b *Builder) getSources() (err error) {
	for i, path := range b.Pack.Sources {
		source := source.Source{
			Hash:   b.Pack.HashSums[i],
			Source: path,
			Output: b.Pack.SourceDir,
		}

		err = source.Get()
		if err != nil {
			return
		}
	}

	return
}

func (b *Builder) build() (err error) {
	path := filepath.Join(string(os.PathSeparator), "tmp",
		fmt.Sprintf("pacur_%s_build", b.id))

	script, err := os.Create(path)
	if err != nil {
		err = &BuilderError{
			errors.Wrap(err, "builder: Failed to create build script"),
		}
		return
	}
	defer func() {
		os.Remove(path)
	}()

	for _, cmd := range b.Pack.Build {
		script.WriteString(cmd + "\n")
	}

	cmd := exec.Command("sh", path)
	cmd.Dir = b.Pack.SourceDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &BuilderError{
			errors.Wrapf(err, "builder: Failed to exec build script"),
		}
		return
	}

	return
}

func (b *Builder) Build() (err error) {
	b.id = utils.RandStr(12)

	err = b.initDirs()
	if err != nil {
		return
	}

	err = b.getSources()
	if err != nil {
		return
	}

	err = b.build()
	if err != nil {
		return
	}

	return
}
