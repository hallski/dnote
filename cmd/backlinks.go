package cmd

import (
	"dnote/mdfiles"
	"os"

	"github.com/spf13/cobra"
)

var blCmd = &cobra.Command{
	Use:   "bl",
	Short: "View backlinks to note",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()
		result := notes.Backlinks(mdfiles.PadID(args[0]))
		ListNoteLinks(result, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(blCmd)
}
