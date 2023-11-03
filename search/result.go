package search

import "dnote"

type Result struct {
	result []*dnote.Note
}

func (sr *Result) ListNotes() []*dnote.Note {
	return sr.result
}
