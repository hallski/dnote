package main

import (
	"dnote"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

type NoteLister interface {
	ListNotes() []*dnote.Note
}

func List(lister NoteLister) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	idColor := color.New(color.FgHiYellow).SprintfFunc()
	tagColor := color.New(color.FgGreen).SprintfFunc()

	for _, note := range lister.ListNotes() {
		fmt.Fprintf(w, "%s\t%s\t%s\n",
			idColor("%d", note.Id),
			note.Title,
			tagColor(strings.Join(note.Tags, ", ")))
	}

	w.Flush()
}
