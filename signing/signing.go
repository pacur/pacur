package signing

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"strings"
)

func GetName() (name string, err error) {
	output, err := utils.ExecOutput("", "gpg", "--list-keys")
	if err != nil {
		return
	}

	for _, line := range strings.Split(output, "\n") {
		if !strings.HasPrefix(line, "uid") {
			continue
		}

		index := strings.Index(line, "]")
		if index == -1 {
			fields := strings.Fields(line)
			if len(fields) < 2 {
				continue
			}

			name = strings.Join(fields[1:], " ")
		} else {
			name = strings.TrimSpace(line[index+1:])
		}

		return
	}

	err = &LookupError{
		errors.New("signing: Failed to find gpg name"),
	}
	return
}

func GetId() (id string, err error) {
	output, err := utils.ExecOutput("", "gpg", "--list-keys")
	if err != nil {
		return
	}

	for _, line := range strings.Split(output, "\n") {
		if !strings.HasPrefix(line, "pub") && !strings.HasPrefix(line, "sub") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		split := strings.Split(fields[1], "/")
		if len(split) < 2 {
			continue
		}

		id = split[1]
	}

	if id == "" {
		err = &LookupError{
			errors.New("signing: Failed to find gpg id"),
		}
	}
	return
}

func ImportKey(path string) (err error) {
	utils.Exec("", "gpg", "--batch", "--allow-secret-key-import",
		"--import", path)
	// TODO err handle already imported
	return
}
