package cmd

import (
	"dnote/core"
	"os"

	"github.com/spf13/cobra"
)

var blCmd = &cobra.Command{
	Use:   "bl",
	Short: "View backlinks to note",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := notes.Backlinks(args[0])
		core.ListNoteLinks(result, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(blCmd)
}
