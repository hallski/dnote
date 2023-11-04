package search

import (
	"dnote"
	"slices"
)

func NewIdsSearch(ids []string, lister dnote.NoteLister) *Result {
	var result []*dnote.Note

	for _, note := range lister.ListNotes() {
		if slices.Contains(ids, note.ID) {
			result = append(result, note)
		}
	}

	return &Result{
		result: result,
	}
}
