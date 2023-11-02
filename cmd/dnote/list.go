package main

import (
	"dnote"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"unicode"

	"github.com/fatih/color"
)

type NoteLister interface {
	ListNotes() []*dnote.Note
}

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

func List(lister NoteLister) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	idColor := color.New(color.FgHiYellow).SprintfFunc()
	tagColor := color.New(color.FgGreen).SprintfFunc()

	for _, note := range lister.ListNotes() {
		fmt.Fprintf(w, "%s\t%s\t%s\n",
			idColor("%d", note.Id),
			ellipticalTruncate(note.Title, 42),
			tagColor(strings.Join(note.Tags, ", ")))
	}

	w.Flush()
}
