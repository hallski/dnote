package search

import (
	"dnote"
	"slices"
)

func NewIdsSearch(ids []int, lister dnote.NoteLister) *SearchResult {
	var result []*dnote.Note

	for _, note := range lister.ListNotes() {
		if slices.Contains(ids, note.Id) {
			result = append(result, note)
		}
	}

	return &SearchResult{
		result: result,
	}
}
