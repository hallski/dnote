package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate notebook",
	Long:  "Migrate notebook to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()
		if err := notes.Migrate(); err != nil {
			fmt.Printf("Failed to migrate notebook: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
