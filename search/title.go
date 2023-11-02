package search

import (
	"dnote"
	"strings"
)

type SearchResult struct {
	result []*dnote.Note
}

func NewTitleSearch(query string, lister dnote.NoteLister) *SearchResult {
	var result []*dnote.Note
	for _, note := range lister.ListNotes() {
		if strings.Contains(strings.ToLower(note.Title), strings.ToLower(query)) {
			result = append(result, note)
		}
	}
	return &SearchResult{
		result,
	}
}

func (sr *SearchResult) ListNotes() []*dnote.Note {
	return sr.result
}
