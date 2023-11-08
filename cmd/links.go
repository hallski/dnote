package cmd

import (
	"dnote/core"
	"dnote/mdfiles"
	"dnote/search"
	"os"

	"github.com/spf13/cobra"
)

var linksCmd = &cobra.Command{
	Use:   "links",
	Short: "Create links for IDs",
	Long:  "Create an index link list for notes with IDs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()

		var ids []string
		for _, id := range os.Args[2:] {
			ids = append(ids, mdfiles.PadID(id))
		}

		result := search.NewIdsSearch(ids, notes)

		core.ListNoteLinks(result, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(linksCmd)
}
