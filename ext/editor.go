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

func GetEditorNewPane(path string, keepFocus bool) *exec.Cmd {
	args := []string{
		"@launch",
		"--cwd",
		"current",
		"--copy-env",
	}

	if keepFocus {
		args = append(args, "--keep-focus")
	}

	editor := GetEditor()
	args = append(args, editor, path)

	cmd := exec.Command("kitten", args...)
	return cmd
}

func EditNote(note *core.Note) error {
	return Exec(GetEditor(), note.Path)
}

func OpenEditor() error {
	return Exec(GetEditor())
}
