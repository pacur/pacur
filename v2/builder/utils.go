package builder

import (
	"github.com/aanno/pacur/v2/utils"
)

func createScript(path string, cmds []string) (err error) {
	data := "set -e\n"
	for _, cmd := range cmds {
		data += cmd + "\n"
	}

	err = utils.CreateWrite(path, data)
	if err != nil {
		return
	}

	return
}

func runScript(path, dir string) (err error) {
	err = utils.Exec(dir, "sh", path)
	if err != nil {
		return
	}

	return
}
