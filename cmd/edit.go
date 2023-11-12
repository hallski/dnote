package cmd

import (
	"fmt"

	"dnote/core"
	"dnote/ext"

	"github.com/spf13/cobra"
)

type NoteFinder interface {
	FindNote(id string) *core.Note
}

func Edit(note *core.Note) error {
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
			note := notes.FindNote(args[0])
			if note == nil {
				return fmt.Errorf("Couldn't find note %s", args[0])
			}
			return Edit(note)
		} else {
			return ext.OpenEditor()
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
