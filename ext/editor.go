package ext

import (
	"dnote/config"
	"dnote/core"
	"os"
	"os/exec"
)

// Interact with $EDITOR

func GetEditor() string {
	cfg := config.GetConfig()

	// If use_environment is true and EDITOR is set, use it
	if cfg.Editor.UseEnvironment {
		if editor := os.Getenv("EDITOR"); editor != "" {
			return editor
		}
	}

	// Use configured editor command if set
	if cfg.Editor.Command != "" {
		return cfg.Editor.Command
	}

	// Fallback to EDITOR environment variable
	return os.Getenv("EDITOR")
}

func GetEditorNewPane(path string, keepFocus bool) *exec.Cmd {
	cfg := config.GetConfig()

	// Start with configured terminal args
	args := make([]string, len(cfg.Editor.TerminalArgs))
	copy(args, cfg.Editor.TerminalArgs)

	if keepFocus {
		args = append(args, "--keep-focus")
	}

	editor := GetEditor()
	args = append(args, editor, path)

	cmd := exec.Command(cfg.Editor.Terminal, args...)
	return cmd
}

func EditNote(note *core.Note) error {
	return Exec(GetEditor(), note.Path)
}

func OpenEditor() error {
	return Exec(GetEditor())
}
