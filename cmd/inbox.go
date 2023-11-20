package cmd

import (
	"dnote/mdfiles"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var inboxCmd = &cobra.Command{
	Use:   "inbox",
	Short: "Add to inbox",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		inbox := notes.GetInbox()
		if inbox == nil {
			return fmt.Errorf("Couldn't find inbox")
		}
		return mdfiles.AddToInbox(inbox.Path, strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(inboxCmd)
}
