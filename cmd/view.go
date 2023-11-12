package cmd

import (
	"fmt"

	"dnote/core"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func View(note *core.Note, random bool) {
	out, err := glamour.Render(note.Content, "dracula")
	if err != nil {
		fmt.Printf("Failed to render note: %s\n", note.ID)
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#55ff55")).
		Underline(true)

	idStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#55ff55"))

	titleStyle := lipgloss.NewStyle().Margin(1, 2)

	var descr string
	if random {
		descr = style.Render("Showing random note:")
	} else {
		descr = style.Render("Showing note:")
	}

	id := idStyle.Render(note.ID)
	title := titleStyle.Render(lipgloss.JoinHorizontal(0, descr, " ", id))

	fmt.Printf("%s%s", title, out)
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View note",
	Long:  "View a note without opening it in editor",
	RunE: func(cmd *cobra.Command, args []string) error {
		var note *core.Note

		random := false
		if len(args) > 0 {
			note = notes.FindNote(args[0])
			if note == nil {
				return fmt.Errorf("Couldn't find note %s", args[0])
			}
		} else {
			random = true
			note = notes.RandomNote()
		}

		View(note, random)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
