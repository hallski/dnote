package search

import "dnote/core"

type Result struct {
	Result []*core.Note
}

func (sr *Result) ListNotes() []*core.Note {
	return sr.Result
}
