package search

import (
	"dnote/core"
	"fmt"
	"log"
	"os/exec"
	"slices"
	"strings"
)

func getUgrepCommand(path, query string) *exec.Cmd {
	searchString := fmt.Sprintf("%s -##nosearch", query)
	args := []string{
		"-l",      // Show matching files
		"-i",      // Ignore case
		"-r",      // Recursive
		"--bool",  // Boolean search
		"--files", // Match in entire file
		"--format=%a%~",
		searchString,
		path,
	}

	ugrepPath, err := exec.LookPath("ugrep")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(ugrepPath, args...)

	return cmd
}

type FullTextCollection interface {
	Path() string
	ListNotes() []*core.Note
}

func NewFullText(query string, collection FullTextCollection) *Result {
	cmd := getUgrepCommand(collection.Path(), query)
	output, _ := cmd.CombinedOutput()

	log.Printf("Searching with cmd: %s\n", cmd)

	files := strings.Split(string(output), "\n")
	var ids []string
	for _, id := range files {
		if !strings.HasSuffix(id, ".md") {
			continue
		}

		ids = append(ids, id[:len(id)-3])
	}
	var notes []*core.Note
	for _, note := range collection.ListNotes() {
		if slices.Contains(ids, note.ID) {
			notes = append(notes, note)
		}
	}
	log.Printf("Result with %d files, resulting in %d notes\n", len(files), len(notes))
	log.Printf("Files: \n%v\n", files)

	slices.SortFunc(notes, func(a, b *core.Note) int {
		return len(b.BackLinks.ListNotes()) - len(a.BackLinks.ListNotes())
	})

	return &Result{query, notes}
}
