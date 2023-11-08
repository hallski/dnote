package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a note file",
	Long:  "Rename a note file and update all links to it",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := notes.Rename(args[0], args[1]); err != nil {
			return fmt.Errorf("Failed to rename %s to %s: %s\n", args[0], args[1], err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)
}
