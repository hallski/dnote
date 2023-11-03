package main

import (
	"dnote"
	"fmt"
	"log"
)

type NoteFinder interface {
	FindNote(id int) *dnote.Note
}

func Open(id int, finder NoteFinder) {
	note := finder.FindNote(id)
	if note == nil {
		fmt.Printf("Couldn't find note %d", id)
		return
	}

	if err := Edit(note); err != nil {
		log.Fatalf("Error while editing file %v", err)
	}
}
