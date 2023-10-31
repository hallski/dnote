package main

import (
	"fmt"
	"log"
	"strconv"

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

func main() {
	argLength := len(os.Args[1:])
	vaultPath := getVaultPath()

	vault, err := LoadVault(vaultPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("main():", vault.NextId())

	cmd := os.Args[1]
	fmt.Println("Command is", cmd)
	if cmd == "edit" {
		if argLength < 2 {
			panic("You must provide a command and a note id")
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Error while editing file %v", err)
		}

		note := vault.GetNote(id)
		if err := Edit(note.Path); err != nil {
			log.Fatalf("Error while editing file %v", err)
		}
	} else if cmd == "new" {
		vault.CreateNote()
	} else if cmd == "ls" {
		fmt.Println("Listing notes")
	} else {
		fmt.Println("No valid command given")
	}
}
