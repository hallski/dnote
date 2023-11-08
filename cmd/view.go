package cmd

import (
	"fmt"

	"dnote/core"

	"github.com/charmbracelet/glamour"
)

func View(note *core.Note) {
	out, err := glamour.Render(note.Content, "dracula")
	if err != nil {
		fmt.Printf("Failed to render note: %s\n", note.ID)
	}

	fmt.Print(out)
}
