package search

import (
	"dnote/core"
	"os/exec"
	"slices"
	"strings"
)

func NewFullText(query string, lister core.NoteLister) *Result {
	cmd := exec.Command("rg", "-l", query)
	output, _ := cmd.CombinedOutput()

	files := strings.Split(string(output), "\n")
	var ids []string
	for _, id := range files {
		if !strings.HasSuffix(id, ".md") {
			continue
		}

		ids = append(ids, id[:len(id)-3])
	}

	var notes []*core.Note
	for _, note := range lister.ListNotes() {
		if slices.Contains(ids, note.ID) {
			notes = append(notes, note)
		}
	}

	return &Result{notes}
}
