package main

import (
	"dnote"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"unicode"

	"github.com/charmbracelet/lipgloss"
)

var bracketStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("4"))
var idStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("11"))

var tagStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("10"))

// https://stackoverflow.com/a/73939904
func ellipticalTruncate(text string, maxLen int) string {
	lastSpaceIx := maxLen
	len := 0
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		len++
		if len > maxLen {
			return text[:lastSpaceIx] + "..."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}

func List(lister dnote.NoteLister, out io.Writer) {
	w := tabwriter.NewWriter(out, 0, 0, 1, ' ', 0)

	for _, note := range lister.ListNotes() {
		fmt.Fprintf(w, "%s%s%s\t%s\t%s\n",
			bracketStyle.Render("["),
			idStyle.Render(fmt.Sprintf("%s", note.ID)),
			bracketStyle.Render("]"),
			ellipticalTruncate(note.Title, 42),
			tagStyle.Render(strings.Join(note.Tags, ", ")))
	}

	w.Flush()
}

func ListNoteLinks(lister dnote.NoteLister, out io.Writer) {
	const linkChars = 4

	for _, note := range lister.ListNotes() {
		truncated := ellipticalTruncate(note.Title, 65)

		padLen := 75 - len([]rune(truncated)) - linkChars - len([]rune(note.ID))
		dots := strings.Repeat(".", padLen)

		fmt.Fprintf(out, "%s%s[[%s]]\n", truncated, dots, note.ID)
	}
}
