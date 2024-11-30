package cmd

import (
	"dnote/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Open TUI",
	Long:  "Open terminal UI for interactiving with notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Getenv("DEBUG")) > 0 {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				fmt.Println("fatal:", err)
				os.Exit(1)
			}
			defer f.Close()
		}

		return tui.Run(getNotesPath(), args[0])
	},
}

func init() {
	rootCmd.AddCommand(browseCmd)
}
