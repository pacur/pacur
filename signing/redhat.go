package signing

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"os/user"
	"path/filepath"
)

func CreateRedhatConf() (err error) {
	name, err := GetName()
	if err != nil {
		return
	}

	data := fmt.Sprintf("%_signature gpg\n%_gpg_name %s\n", name)

	usr, err := user.Current()
	if err != nil {
		err = &LookupError{
			errors.Wrap(err, "signing: Failed to get current user"),
		}
		return
	}

	err = utils.CreateWrite(filepath.Join(usr.HomeDir, ".rpmmacros"), data)
	if err != nil {
		return
	}

	return
}

func SignRedhat(dir string) (err error) {
	err = CreateRedhatConf()
	if err != nil {
		return
	}

	pkgs, err := utils.FindExt(dir, ".rpm")
	if err != nil {
		return
	}

	for _, pkg := range pkgs {
		err = utils.Exec("", "rpm", "--resign", pkg)
		if err != nil {
			return
		}
	}

	return
}
