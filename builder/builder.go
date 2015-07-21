package builder

import (
	"github.com/pacur/pacur/pack"
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
