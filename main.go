package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"strconv"
	"text/tabwriter"

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

func Edit(path string) error {
	return Execute("nvim", path)
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

func List(vault *Vault) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	yellow := color.New(color.FgHiYellow).SprintfFunc()

	for _, note := range vault.Notes {
		fmt.Fprintf(w, "%s\t%s\n", yellow("%d", note.Id), note.Title)
	}

	w.Flush()
}

func main() {
	argLength := len(os.Args[1:])
	vaultPath := getVaultPath()

	vault, err := LoadVault(vaultPath)
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

		note := vault.GetNote(id)
		if note == nil {
			fmt.Printf("Couldn't find note %d", id)
			return
		}

		if err := Edit(note.Path); err != nil {
			log.Fatalf("Error while editing file %v", err)
		}
	} else if cmd == "new" {
		vault.CreateNote()
	} else if cmd == "ls" {
		List(vault)
	} else {
		fmt.Println("No valid command given")
	}
}
