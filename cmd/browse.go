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
		var openId = ""
		if len(args) > 0 {
			openId = mdfiles.PadID(args[0])
		}
		return tui.Run(notes, openId)
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}
