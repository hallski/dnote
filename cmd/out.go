package cmd

import (
	"dnote/mdfiles"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var outCmd = &cobra.Command{
	Use:   "out",
	Short: "View outgoing links from note",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		note := notes.FindNote(mdfiles.PadID(args[0]))
		if note != nil {
			fmt.Println(strings.Join(note.Links, " "))
		}
	},
}

func init() {
	rootCmd.AddCommand(outCmd)
}
