package main

import (
	"dnote"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode"

	"github.com/fatih/color"
)

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

func List(lister dnote.NoteLister) {
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

func ListNoteLinks(lister dnote.NoteLister) {
	const linkChars = 4

	for _, note := range lister.ListNotes() {
		truncated := ellipticalTruncate(note.Title, 60)
		strId := strconv.Itoa(note.Id)

		padLen := 70 - len([]rune(truncated)) - linkChars - len([]rune(strId))

		fmt.Printf("%s%s[[%s]]\n", truncated, strings.Repeat(".", padLen), strId)
	}
}
