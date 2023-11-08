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

func List(lister core.NoteLister, out io.Writer) {
	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', 0)

	for _, note := range lister.ListNotes() {
		fmt.Fprintf(w, "%s%s%s\t%s\t%s\n",
			bracketStyle.Render("["),
			idStyle.Render(fmt.Sprintf("%s", note.ID)),
			bracketStyle.Render("]"),
			core.EllipticalTruncate(note.Title, 42),
			tagStyle.Render(strings.Join(note.Tags, ", ")))
	}

	w.Flush()
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all notes",
	Long:  "List all files together with ID",
	Run: func(cmd *cobra.Command, args []string) {
		List(notes, os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
