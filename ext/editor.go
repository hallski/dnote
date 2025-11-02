package ext

import (
	"fmt"
	"os"
	"os/exec"

	"dnote/config"
	"dnote/core"
)

// Interact with $EDITOR

func GetEditor() string {
	// If use_environment is true and EDITOR is set, use it
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	} else {
		fmt.Errorf("No editor defined")
	}

	return ""
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
