package main

import (
	"dnote"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

func List(storage dnote.NoteStorage) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	yellow := color.New(color.FgHiYellow).SprintfFunc()

	for _, note := range storage.AllNotes() {
		fmt.Fprintf(w, "%s\t%s\n", yellow("%d", note.Id), note.Title)
	}

	w.Flush()
}
