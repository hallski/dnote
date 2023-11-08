package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create and open new note",
	Long:  "Creates a new note with the next available ID and opens it in editor",
	RunE: func(cmd *cobra.Command, _ []string) error {
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}

		note, err := notes.CreateNote(title)
		if err != nil {
			return err
		}

		return Edit(note)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "Title of the new note")
}
