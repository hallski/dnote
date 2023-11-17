package ext

import (
	"dnote/core"
	"os"
	"os/exec"
)

// Interact with $EDITOR

func GetEditor() string {
	return os.Getenv("EDITOR")
}

func GetEditorNewPane(path string) *exec.Cmd {
	editor := GetEditor()
	cmd := exec.Command("kitten", "@launch", editor, path)
	return cmd
}

func EditNote(note *core.Note) error {
	return Exec(GetEditor(), note.Path)
}

func OpenEditor() error {
	return Exec(GetEditor())
}
