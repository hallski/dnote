package cmd

import (
	"dnote/search"
	"os"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search note titles",
	Long:  "Search note titles for strings containing query and list as index",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result := search.NewTitleSearch(os.Args[2], notes)
		List(result, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
