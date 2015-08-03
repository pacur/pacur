package signing

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"strings"
)

func GetName() (name string, err error) {
	output, err := utils.ExecOutput("", "gpg", "-K")
	if err != nil {
		return
	}

	for _, line := range strings.Split(output, "\n") {
		if !strings.HasPrefix(line, "uid") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		name = strings.Join(fields[1:], " ")
		return
	}

	err = &LookupError{
		errors.New("signing: Failed to find gpg name"),
	}
	return
}

func ImportKey(path string) (err error) {
	err = utils.Exec("", "gpg", "--allow-secret-key-import", "--import", path)
	if err != nil {
		return
	}

	return
}
