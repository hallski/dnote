package cmd

import (
	"fmt"

	"dnote/core"
	"dnote/mdfiles"

	"github.com/spf13/cobra"
)

type NoteFinder interface {
	FindNote(id string) *core.Note
}

func Open(id string, finder NoteFinder) error {
	note := finder.FindNote(id)
	if note == nil {
		return fmt.Errorf("Couldn't find note %s", id)
	}

	if err := Edit(note); err != nil {
		return err
	}

	return nil
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a note",
	Long:  "Opens note with ID in Vim",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			return Open(mdfiles.PadID(args[0]), notes)
		} else {
			return OpenEditor()
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
