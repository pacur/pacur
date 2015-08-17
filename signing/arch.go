package signing

import (
	"github.com/pacur/pacur/utils"
)

func SignArch(dir string) (err error) {
	err = CreateRedhatConf()
	if err != nil {
		return
	}

	pkgs, err := utils.FindExt(dir, ".rpm")
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
