package ext

import (
	"os"
	"os/exec"
)

func Exec(command string, arg ...string) error {
	editorPath, err := exec.LookPath(command)
	if err != nil {
		return err
	}

	cmd := exec.Command(editorPath, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
