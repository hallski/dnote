package cmd

import (
	"os"
	"os/exec"

	"dnote/core"
)

const editor = "nvim"

func RunCmd(command string, arg ...string) error {
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

func Edit(note *core.Note) error {
	return RunCmd(editor, note.Path)
}

func OpenEditor() error {
	return RunCmd(editor)
}
