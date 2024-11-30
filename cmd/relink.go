package cmd

import (
	"bufio"
	"dnote/core"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var relinkCmd = &cobra.Command{
	Use:   "relink",
	Short: "Recreate links from stdin",
	RunE: func(cmd *cobra.Command, args []string) error {
		notes := loadNotes()
		reader := bufio.NewReader(os.Stdin)

		for {
			s, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			links := core.ExtractLinks(s)
			if len(links) == 0 {
				fmt.Print(s)
			}

			result := notes.GetIds(links...)
			core.ListNoteLinks(result, os.Stdout)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(relinkCmd)
}
