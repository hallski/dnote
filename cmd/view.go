package cmd

import (
	"fmt"

	"dnote/core"
	"dnote/mdfiles"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

func View(note *core.Note) {
	out, err := glamour.Render(note.Content, "dracula")
	if err != nil {
		fmt.Printf("Failed to render note: %s\n", note.ID)
	}

	fmt.Print(out)
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View note",
	Long:  "View a note without opening it in editor",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()
		note := notes.FindNote(mdfiles.PadID(args[0]))
		View(note)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
