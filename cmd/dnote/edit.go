package main

import (
	"dnote"
	"os"
	"os/exec"
)

func Execute(command string, arg ...string) error {
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

func Edit(note *dnote.Note) error {
	return Execute("nvim", note.Path)
}
