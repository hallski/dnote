package cmd

import (
	"dnote/mdfiles"
	"dnote/tui"

	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Open TUI",
	Long:  "Open terminal UI for interactiving with notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		var openId = "1"
		if len(args) > 0 {
			note := notes.FindNote(args[0])
			if note != nil {
				openId = note.ID
			}
		}

		return tui.Run(notes, mdfiles.PadID(openId))
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}
