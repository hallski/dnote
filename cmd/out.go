package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var outCmd = &cobra.Command{
	Use:   "out",
	Short: "View outgoing links from note",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		note := notes.FindNote(args[0])
		if note != nil {
			return fmt.Errorf(strings.Join(note.Links, " "))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(outCmd)
}
