package search

import (
	"dnote/core"
	"strings"
)

func NewTitleSearch(query string, lister core.NoteLister) *Result {
	var result []*core.Note
	for _, note := range lister.ListNotes() {
		if strings.Contains(strings.ToLower(note.Title), strings.ToLower(query)) {
			result = append(result, note)
		}
	}
	return &Result{
		result,
	}
}
