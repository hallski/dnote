package search

import (
	"slices"

	"dnote/core"
)

func NewIdsSearch(ids []string, lister core.NoteLister) *Result {
	var result []*core.Note

	for _, note := range lister.ListNotes() {
		if slices.Contains(ids, note.ID) {
			result = append(result, note)
		}
	}

	return &Result{
		result: result,
	}
}
