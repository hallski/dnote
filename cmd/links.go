package cmd

import (
	"dnote/core"
	"dnote/mdfiles"
	"dnote/search"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var linksCmd = &cobra.Command{
	Use:   "links",
	Short: "Create links for IDs",
	Long:  "Create an index link list for notes with IDs",
	RunE: func(cmd *cobra.Command, args []string) error {
		useStored, err := cmd.Flags().GetBool("stored")
		if err != nil {
			return err
		}

		var result core.NoteLister
		if useStored {
			result, err = notes.GetCollection()
			if err != nil {
				return err
			}
		} else {
			if len(args) < 1 {
				return fmt.Errorf("Need at least one ID")
			}
			var ids []string
			for _, id := range args {
				if id == "last" {
					last := notes.LastNote()
					id = last.ID
				}
				ids = append(ids, mdfiles.PadID(id))
			}

			result = search.NewIdsSearch(ids, notes)
		}

		core.ListNoteLinks(result, os.Stdout)

		return nil
	},
}

func init() {
	linksCmd.PersistentFlags().BoolP("stored", "s", false, "Use stored values")

	rootCmd.AddCommand(linksCmd)
}
