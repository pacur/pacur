package builder

import (
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/source"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type Builder struct {
	srcDir string
	Pack   *pack.Pack
}

func (b *Builder) initDirs() (err error) {
	b.srcDir = filepath.Join(b.Pack.Root, "src")
	err = utils.ExistsMakeDir(b.srcDir)
	if err != nil {
		return
	}

	return
}

func (b *Builder) getSources() (err error) {
	for _, path := range b.Pack.Sources {
		source := source.Source{
			Path:   path,
			Output: b.srcDir,
		}

		err = source.Get()
		if err != nil {
			return
		}
	}

	return
}

func (b *Builder) Build() (err error) {
	err = b.initDirs()
	if err != nil {
		return
	}

	err = b.getSources()
	if err != nil {
		return
	}

	return
}
