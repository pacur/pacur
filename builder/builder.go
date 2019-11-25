package builder

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/source"
	"github.com/pacur/pacur/utils"
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
			Root:   b.Pack.Root,
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
	defer os.Remove(path)

	err = createScript(path, b.Pack.Build)
	if err != nil {
		return
	}

	err = runScript(path, b.Pack.SourceDir)
	if err != nil {
		return
	}

	return
}

func (b *Builder) pkg() (err error) {
	path := filepath.Join(string(os.PathSeparator), "tmp",
		fmt.Sprintf("pacur_%s_package", b.id))
	defer os.Remove(path)

	err = createScript(path, b.Pack.Package)
	if err != nil {
		return
	}

	err = runScript(path, b.Pack.SourceDir)
	if err != nil {
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

	err = b.pkg()
	if err != nil {
		return
	}

	return
}
