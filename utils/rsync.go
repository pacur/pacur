package utils

import (
	"github.com/dropbox/godropbox/errors"
	"os"
	"os/exec"
)

func Rsync(source, dest string) (err error) {
	cmd := exec.Command("rsync", "-a", "-A", "-X",
		source+string(os.PathSeparator),
		dest+string(os.PathSeparator))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &CopyError{
			errors.Wrapf(err, "utils: Failed to rsync '%s' to '%s'", source,
				dest),
		}
		return
	}

	return
}

func RsyncExt(source, dest, ext string) (err error) {
	cmd := exec.Command("rsync", "-a", "-A", "-X",
		"--include", "*."+ext, "--exclude", "*",
		source+string(os.PathSeparator),
		dest+string(os.PathSeparator))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		err = &CopyError{
			errors.Wrapf(err, "utils: Failed to rsync '%s' to '%s'", source,
				dest),
		}
		return
	}

	return
}
