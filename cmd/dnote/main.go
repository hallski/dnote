package main

import (
	"dnote/mdfiles"
	"dnote/search"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getNotesPath() string {
	path := os.Getenv("DNOTES")
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			path = "."
		}
	}

	return path
}

func main() {
	argLen := len(os.Args[1:])
	path := getNotesPath()

	notes, err := mdfiles.Load(path)
	if err != nil {
		panic(err)
	}

	cmd := os.Args[1]
	if cmd == "open" {
		if argLen < 2 {
			panic("You must provide a command and a note id")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error while editing file %v", err)
		}

		Open(id, notes)
	} else if cmd == "new" {
		note, err := notes.CreateNote()
		if err != nil {
			log.Fatalf("Couldn't create new note: %s", err)
		}

		if err := Edit(note); err != nil {
			log.Fatalf("Error opening new note %s", err)
		}
	} else if cmd == "ls" {
		List(notes, os.Stdout)
	} else if cmd == "search" {
		if argLen < 2 {
			panic("No search query")
		}

		result := search.NewTitleSearch(os.Args[2], notes)
		ListNoteLinks(result, os.Stdout)
	} else if cmd == "ids" {
		if argLen < 2 {
			panic("Need to give a list of ids")
		}

		var ids []int
		for _, strId := range os.Args[2:] {
			id, err := strconv.Atoi(strId)
			if err != nil {
				continue
			}

			ids = append(ids, id)
		}

		result := search.NewIdsSearch(ids, notes)
		ListNoteLinks(result, os.Stdout)
	} else if cmd == "version" {
		fmt.Println("Version 0.2")
	} else {
		fmt.Println("No valid command given")
	}
}
