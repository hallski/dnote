package main

import (
	"dnote"
	"dnote/mdfiles"
	"fmt"
	"log"
	"strconv"
	"text/tabwriter"

	"github.com/fatih/color"

	"os"
	"os/exec"
)

func Execute(command string, arg ...string) error {
	editorPath, err := exec.LookPath(command)
	if err != nil {
		return err
	}

	cmd := exec.Command(editorPath, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func Edit(note *dnote.Note) error {
	return Execute("nvim", note.Path)
}

func SearchAndOpenByTitle(vaultPath string) {
	return
	// List all files
	// Reach title for each file
	// Map to "Id   Title" and pipe to FZF
	// Get Id from selected line
	// Edit(Id)
}

func getVaultPath() string {
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

func List(storage dnote.NoteStorage) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	yellow := color.New(color.FgHiYellow).SprintfFunc()

	for _, note := range storage.AllNotes() {
		fmt.Fprintf(w, "%s\t%s\n", yellow("%d", note.Id), note.Title)
	}

	w.Flush()
}

func main() {
	argLength := len(os.Args[1:])
	vaultPath := getVaultPath()

	storage, err := mdfiles.Load(vaultPath)
	if err != nil {
		panic(err)
	}

	cmd := os.Args[1]
	fmt.Println("Command is", cmd)
	if cmd == "open" {
		if argLength < 2 {
			panic("You must provide a command and a note id")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error while editing file %v", err)
		}

		note := storage.FindNote(id)
		if note == nil {
			fmt.Printf("Couldn't find note %d", id)
			return
		}

		if err := Edit(note); err != nil {
			log.Fatalf("Error while editing file %v", err)
		}
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
	} else {
		fmt.Println("No valid command given")
	}
}
