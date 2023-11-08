package search

import "dnote/core"

type Result struct {
	result []*core.Note
}

func (sr *Result) ListNotes() []*core.Note {
	return sr.result
}
