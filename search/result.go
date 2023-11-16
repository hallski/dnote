package search

import "dnote/core"

type Result struct {
	Query  string
	Result []*core.Note
}

func (sr *Result) ListNotes() []*core.Note {
	return sr.Result
}
