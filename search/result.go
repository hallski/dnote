package search

import "dnote"

type SearchResult struct {
	result []*dnote.Note
}

func (sr *SearchResult) ListNotes() []*dnote.Note {
	return sr.result
}
