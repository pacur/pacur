package mirror

import (
	"github.com/dropbox/godropbox/errors"
	"os"
	"os/exec"
)

type Mirror struct {
	Root string
}

func (m *Mirror) createRedhat() (err error) {
	cmd := exec.Command("createrepo", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = m.Root

	err = cmd.Run()
	if err != nil {
		err = &BuildError{
			errors.Wrapf(err, "mirror: Failed to create redhat repo '%s'",
				m.Root),
		}
		return
	}

	return
}

func (m *Mirror) Create(typ string) (err error) {
	switch typ {
	case "redhat":
		err = m.createRedhat()
	default:
		err = &UnknownType{
			errors.Newf("mirror: Unknown type '%s'", typ),
		}
	}

	return
}
