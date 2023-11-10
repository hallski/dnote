package ext

import (
	"dnote/core"
	"os"
)

// Interact with $EDITOR

func GetEditor() string {
	return os.Getenv("EDITOR")
}

func EditNote(note *core.Note) error {
	return Exec(GetEditor(), note.Path)
}

func OpenEditor() error {
	return Exec(GetEditor())
}
