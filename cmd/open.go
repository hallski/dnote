package cmd

import (
	"fmt"
	"log"

	"dnote/core"
)

type NoteFinder interface {
	FindNote(id string) *core.Note
}

func Open(id string, finder NoteFinder) {
	note := finder.FindNote(id)
	if note == nil {
		fmt.Printf("Couldn't find note %s", id)
		return
	}

	if err := Edit(note); err != nil {
		log.Fatalf("Error while editing file %v", err)
	}
}
