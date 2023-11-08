package cmd

import (
	"dnote/core"
	"dnote/mdfiles"
	"os"

	"github.com/spf13/cobra"
)

var blCmd = &cobra.Command{
	Use:   "bl",
	Short: "View backlinks to note",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := notes.Backlinks(mdfiles.PadID(args[0]))
		core.ListNoteLinks(result, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(blCmd)
}
