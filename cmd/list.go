package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"dnote/core"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var bracketStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("8"))
var idStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("4"))
var tagStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10"))

func List(lister core.NoteLister, out io.Writer, showTags bool) {
	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', 0)

	var titleLen = 80
	if showTags {
		titleLen = 42
	}

	for _, note := range lister.ListNotes() {
		fmt.Fprintf(w, "%s%s%s\t%s",
			bracketStyle.Render("["),
			idStyle.Render(fmt.Sprintf("%s", note.ID)),
			bracketStyle.Render("]"),
			core.EllipticalTruncate(note.Title, titleLen))

		if showTags {
			fmt.Fprintf(w, "\t%s", tagStyle.Render(strings.Join(note.Tags, ", ")))
		}

		fmt.Fprint(w, "\n")
	}

	w.Flush()
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all notes",
	Long:  "List all files together with ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		showTags, err := cmd.Flags().GetBool("tags")
		if err != nil {
			return err
		}

		List(notes, os.Stdout, showTags)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.PersistentFlags().BoolP("tags", "t", false, "List tags as well")
}
