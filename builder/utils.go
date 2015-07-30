package builder

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"os"
	"os/exec"
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
	cmd := exec.Command("sh", path)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &ScriptError{
			errors.Wrapf(err, "builder: Failed to exec script"),
		}
		return
	}

	return
}
