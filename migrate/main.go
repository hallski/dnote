package main

import (
	"bufio"
	"bytes"
	"dnote/core"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	vaultPath := os.Getenv("DNOTES")

	Migrate(vaultPath)
}

type noteEntry struct {
	oldPath     string
	base        string
	oldId       string
	title       string
	timestamp   time.Time
	newId       string
	newPath     string
	newFilename string
}

func MigrateNote(directory string, note *noteEntry) error {
	err := os.Rename(note.oldPath, note.newPath)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(note.newPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	fmt.Fprintf(w, "\n\n[:created]: _ \"%s\"\n", note.timestamp.Format("2006-01-02 15:04"))
	w.Flush()

	f.Close()

	err = core.UpdateLinks(directory, note.oldId, note.newId, note.title)
	if err != nil {
		panic(err)
	}

	return nil
}

// Migrate from old format to new format
func Migrate(directory string) error {

	newId := 1

	var noteFiles []*noteEntry

	filepath.WalkDir(directory, func(path string, _ fs.DirEntry, e error) error {
		if e != nil {
			return e
		}

		if filepath.Ext(path) == ".md" {
			base := filepath.Base(path)
			splits := strings.Split(base, " ")
			oldId := splits[0]
			title := strings.TrimSuffix(strings.Join(splits[1:], " "), filepath.Ext(path))
			timestamp, err := core.ParseZkID(oldId)
			if err != nil {
				panic(err)
			}

			note := &noteEntry{
				oldPath:   path,
				base:      base,
				oldId:     oldId,
				title:     title,
				timestamp: timestamp,
			}
			//	fmt.Printf("%+v", note)

			noteFiles = append(noteFiles, note)
		}

		return nil
	})

	sort.Slice(noteFiles, func(i, j int) bool {
		return noteFiles[i].oldId < noteFiles[j].oldId
	})

	// Assign new Ids
	for _, note := range noteFiles {
		note.newId = strconv.Itoa(newId)
		note.newFilename = fmt.Sprintf("%s.md", note.newId)
		note.newPath = filepath.Join(directory, note.newFilename)
		newId++
		//fmt.Printf("%+v\n\n", note)
	}

	var buf bytes.Buffer
	for _, note := range noteFiles {
		MigrateNote(directory, note)
		fmt.Fprintf(&buf, "%s -> %s %s\n", note.oldId, note.newId, note.title)
	}

	os.WriteFile(filepath.Join(directory, "migration.log"), []byte(buf.String()), 0644)

	return nil
}
