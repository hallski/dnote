package cmd

import (
	"fmt"
	"log"

	"dnote/core"
	"dnote/mdfiles"

	"github.com/spf13/cobra"
)

type NoteFinder interface {
	FindNote(id string) *core.Note
}

func Open(id string, finder NoteFinder) {
	note := finder.FindNote(id)
	if note == nil {
		fmt.Printf("Couldn't find note %s", id)
		return
	}

	if err := Edit(note); err != nil {
		log.Fatalf("Error while editing file %v", err)
	}
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a note",
	Long:  "Opens note with ID in Vim",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			Open(mdfiles.PadID(args[0]), notes)
		} else {
			OpenEditor()
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
