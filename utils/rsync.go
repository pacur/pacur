package utils

import (
	"os"
	"os/exec"

	"github.com/pritunl/tools/errors"
)

func Rsync(source, dest string) (err error) {
	cmd := exec.Command("rsync", "-a", "-A",
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
	cmd := exec.Command("rsync", "-a", "-A",
		"--include", "*"+ext, "--exclude", "*",
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

func RsyncRelExt(source, dest, release, ext string) (err error) {
	cmd := exec.Command("rsync", "-a", "-A",
		"--include", "*"+release+"*"+ext, "--exclude", "*",
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

func RsyncMatch(source, dest, match string) (err error) {
	cmd := exec.Command("rsync", "-a", "-A",
		"--include", "*"+match+"*", "--exclude", "*",
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
