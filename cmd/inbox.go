package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var inboxCmd = &cobra.Command{
	Use:   "inbox",
	Short: "Add to inbox",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return notes.AddToInbox(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(inboxCmd)
}
