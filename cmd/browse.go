package cmd

import (
	"dnote/tui"

	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Open TUI",
	Long:  "Open terminal UI for interactiving with notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		return tui.Run(notes)
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}
