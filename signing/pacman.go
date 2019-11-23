package signing

import (
	"github.com/m0rf30/pacur/utils"
)

func SignPacman(dir string) (err error) {
	pkgs, err := utils.FindExt(dir, ".pkg.tar.xz")
	if err != nil {
		return
	}

	name, err := GetName()
	if err != nil {
		return
	}

	for _, pkg := range pkgs {
		err = utils.Exec(dir, "gpg",
			"--detach-sign",
			"-u", name,
			"--no-armor",
			pkg)
		if err != nil {
			return
		}
	}

	return
}
