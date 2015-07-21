package builder

import (
	"github.com/pacur/pacur/pack"
	"github.com/pacur/pacur/source"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type Builder struct {
	srcDir string
	pac    *pack.Pack
}

func (b *Builder) initDirs() (err error) {
	b.srcDir = filepath.Join(b.pac.Root, "src")
	err = utils.ExistsMakeDir(b.srcDir)
	if err != nil {
		return
	}

	return
}

func (b *Builder) getSources() (err error) {
	for _, path := range b.pac.Sources {
		source := source.Source{
			Path: path,
			Output: b.srcDir,
		}

		err = source.Get()
		if err != nil {
			return
		}
	}

	return
}
