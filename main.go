package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	// "dnote/shell/tui"

	"os"
	"os/exec"
	"path/filepath"
)

const pagesPath = "pages"

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

func Edit(vaultPath, noteId string) error {
	fileName := fmt.Sprintf("%v.md", noteId)
	path := filepath.Join(vaultPath, fileName)

	return Execute("nvim", path)
}

func getFileId(fileName string) (int, error) {
	base, ext, found := strings.Cut(fileName, ".")
	if !found || ext != "md" {
		return -1, fmt.Errorf("Filename not following convention of id.md: %s",
			fileName)
	}

	nr, err := strconv.Atoi(base)
	if err != nil {
		return -1, fmt.Errorf("Filename not following convention of id.md: %s",
			fileName)
	}

	return nr, nil
}

func getFiles(vaultPath string) ([]string, error) {
	files, err := os.ReadDir(vaultPath)
	if err != nil {
		return []string{}, fmt.Errorf("Failed to open vault: %s", err)
	}

	var fileNames []string
	for _, fileEntry := range files {
		fileNames = append(fileNames, fileEntry.Name())
	}

	return fileNames, nil
}

func getLastId(vaultPath string) (int, error) {
	files, err := getFiles(vaultPath)
	if err != nil {
		return -1, err
	}

	var highest = 0
	for _, fileName := range files {
		nr, err := getFileId(fileName)
		if err != nil {
			log.Println("Warning", err)
			continue
		}

		if nr > highest {
			highest = nr
		}
	}

	return highest, nil
}

func AddTimestampToNote(dir, id string, timestamp time.Time) error {
	filename := fmt.Sprintf("%s.md", id)
	path := filepath.Join(dir, filename)

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Failed to append timestamp: %s", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "\n\n[:created]: _ \"%s\"\n",
		timestamp.Format("2006-01-02 15:04"))

	return nil
}

func Create(vaultPath string) error {
	// Find the last ID used
	highest, err := getLastId(vaultPath)
	if err != nil {
		return err
	}

	// Increment the ID
	highest++

	newId := strconv.Itoa(highest)

	AddTimestampToNote(vaultPath, newId, time.Now())

	err = Edit(vaultPath, newId)
	if err != nil {
		return err
	}

	return nil
}

func SearchAndOpenByTitle(vaultPath string) {
	_, err := getFiles(vaultPath)
	if err != nil {
		log.Println("Error:", err)
		return
	}

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

		vaultPath = filepath.Join(vaultPath, "example")
	}

	return vaultPath
}

func main() {
	argLength := len(os.Args[1:])
	vaultPath := getVaultPath()

	cmd := os.Args[1]
	fmt.Println("Command is", cmd)
	if cmd == "edit" {
		if argLength < 2 {
			panic("You must provide a command and a note id")
		}
		if err := Edit(vaultPath, os.Args[2]); err != nil {
			log.Fatalf("Error while editing file %v", err)
		}
	} else if cmd == "new" {
		Create(vaultPath)
		// }
		// else if cmd == "tui" {
		// 	fmt.Printf("You have %v arguments\n", argLength)
		//
		// 	tui.StartTui(os.Args[2])
	} else {
		fmt.Println("No valid command given")
	}
}
