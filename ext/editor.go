package ext

import (
	"os"
	"os/exec"

	"dnote/config"
	"dnote/core"
)

// Interact with $EDITOR

func GetEditor() string {
	return os.Getenv("EDITOR")
}

func GetEditorInteractive(path string) *exec.Cmd {
	cfg := config.GetConfig()

	// Start with configured terminal args
	args := make([]string, len(cfg.Editor.Args))
	copy(args, cfg.Editor.Args)

	//		editor := GetEditor()
	args = append(args, path)

	cmd := exec.Command(cfg.Editor.Command, args...)
	//	cmd := exec.Command(cfg.Editor.Terminal, args...)
	//	cmd := exec.Command(cfg.Editor.Terminal, args...)
	return cmd
}

func EditNote(note *core.Note) error {
	return Exec(GetEditor(), note.Path)
}

func OpenEditor() error {
	return Exec(GetEditor())
}
