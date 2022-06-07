package u

import (
	"os/exec"
)

func Exec(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}
