package signing

import (
	"fmt"
	"github.com/pacur/pacur/utils"
	"path/filepath"
)

type GenKey struct {
	Root  string
	Name  string
	Email string
}

func (g *GenKey) createConf() (path string, err error) {
	err = utils.MkdirAll(g.Root)
	if err != nil {
		return
	}

	path = filepath.Join(g.Root, "genkey")

	err = utils.CreateWrite(path, "Key-Type: 1\n"+
		"Key-Length: 4096\n"+
		"Subkey-Type: 1\n"+
		"Subkey-Length: 4096\n"+
		fmt.Sprintf("Name-Real: %s\n", g.Name)+
		fmt.Sprintf("Name-Email: %s\n", g.Email)+
		"Expire-Date: 0\n")
	if err != nil {
		return
	}

	return
}

func (g *GenKey) Generate() (err error) {
	confPath, err := g.createConf()
	if err != nil {
		return
	}
	defer utils.Remove(confPath)

	err = utils.Exec(g.Root, "gpg", "--batch", "--gen-key", confPath)
	if err != nil {
		return
	}

	return
}

func (g *GenKey) Export() (err error) {
	id, err := GetId()
	if err != nil {
		return
	}

	data, err := utils.ExecOutput(g.Root,
		"gpg", "-a", "--export-secret-keys", id)
	if err != nil {
		return
	}

	err = utils.CreateWrite(filepath.Join(g.Root, "sign.key"), data)
	if err != nil {
		return
	}

	return
}
