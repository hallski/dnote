package search

import (
	"dnote/core"
	"dnote/mdfiles"
	"os/exec"
	"slices"
	"strings"
)

func getUgrepCommand(path, query string) *exec.Cmd {
	args := []string{
		"-l",
		"-i",
		"--bool",
		"--format=%a%~",
		query,
		path,
	}

	ugrepPath, err := exec.LookPath("ugrep")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(ugrepPath, args...)

	return cmd
}

func NewFullText(query string, notebook *mdfiles.MdDirectory) *Result {
	cmd := getUgrepCommand(notebook.Path, query)
	output, _ := cmd.CombinedOutput()

	files := strings.Split(string(output), "\n")
	var ids []string
	for _, id := range files {
		if !strings.HasSuffix(id, ".md") {
			continue
		}

		ids = append(ids, id[:len(id)-3])
	}

	var notes []*core.Note
	for _, note := range notebook.ListNotes() {
		if slices.Contains(ids, note.ID) {
			notes = append(notes, note)
		}
	}

	return &Result{query, notes}
}
