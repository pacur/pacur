package builder

import (
	"github.com/dropbox/godropbox/errors"
	"os"
	"os/exec"
)

func createScript(path string, cmds []string) (err error) {
	script, err := os.Create(path)
	if err != nil {
		err = &ScriptError{
			errors.Wrapf(err, "builder: Failed to create script '%s'", path),
		}
		return
	}
	defer script.Close()

	_, err = script.WriteString("set -e\n")
	if err != nil {
		err = &ScriptError{
			errors.Wrapf(err, "builder: Failed to write script '%s'", path),
		}
		return
	}

	for _, cmd := range cmds {
		_, err = script.WriteString(cmd + "\n")
		if err != nil {
			err = &ScriptError{
				errors.Wrapf(err, "builder: Failed to write script '%s'",
					path),
			}
			return
		}
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
