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
	vaultPath := os.Getenv("DNOTES")
	if vaultPath == "" {
		var err error
		vaultPath, err = os.Getwd()
		if err != nil {
			vaultPath = "."
		}
	}

	return vaultPath
}

func main() {
	argLength := len(os.Args[1:])
	path := getNotesPath()

	storage, err := mdfiles.Load(path)
	if err != nil {
		panic(err)
	}

	cmd := os.Args[1]
	if cmd == "open" {
		if argLength < 2 {
			panic("You must provide a command and a note id")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error while editing file %v", err)
		}

		Open(storage, id)
	} else if cmd == "new" {
		note, err := storage.CreateNote()
		if err != nil {
			log.Fatalf("Couldn't create new note: %s", err)
		}

		if err := Edit(note); err != nil {
			log.Fatalf("Error opening new note %s", err)
		}
	} else if cmd == "ls" {
		List(storage)
	} else if cmd == "search" {
		if argLength < 2 {
			panic("No search query")
		}

		result := search.NewTitleSearch(os.Args[2], storage)
		ListNoteLinks(result)
	} else if cmd == "version" {
		fmt.Println("Version 0.2")
	} else {
		fmt.Println("No valid command given")
	}
}
