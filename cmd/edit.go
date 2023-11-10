package cmd

import (
	"fmt"

	"dnote/core"
	"dnote/ext"
	"dnote/mdfiles"

	"github.com/spf13/cobra"
)

type NoteFinder interface {
	FindNote(id string) *core.Note
}

func Edit(id string, finder NoteFinder) error {
	note := finder.FindNote(id)
	if note == nil {
		return fmt.Errorf("Couldn't find note %s", id)
	}

	if err := ext.EditNote(note); err != nil {
		return err
	}

	return nil
}

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"open"},
	Short:   "Edit a note",
	Long:    "Opens note with ID in $EDITOR",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 {
			return Edit(mdfiles.PadID(args[0]), notes)
		} else {
			return ext.OpenEditor()
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
