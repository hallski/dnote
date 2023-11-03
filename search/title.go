package search

import (
	"dnote"
	"strings"
)

func NewTitleSearch(query string, lister dnote.NoteLister) *Result {
	var result []*dnote.Note
	for _, note := range lister.ListNotes() {
		if strings.Contains(strings.ToLower(note.Title), strings.ToLower(query)) {
			result = append(result, note)
		}
	}
	return &Result{
		result,
	}
}
