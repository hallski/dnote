package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate notebook",
	Long:  "Migrate notebook to latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := notes.Migrate(); err != nil {
			return fmt.Errorf("Failed to migrate notebook: %s\n", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
