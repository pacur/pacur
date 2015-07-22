package builder

import (
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/source"
	"github.com/pacur/pacur/utils"
)

type Builder struct {
	Pack *pack.Pack
}

func (b *Builder) initDirs() (err error) {
	err = utils.ExistsMakeDir(b.Pack.SourceDir)
	if err != nil {
		return
	}

	return
}

func (b *Builder) getSources() (err error) {
	for _, path := range b.Pack.Sources {
		source := source.Source{
			Path:   path,
			Output: b.Pack.SourceDir,
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
