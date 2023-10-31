package main

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"sort"
)

type Vault struct {
	Path  string
	Notes []*Note
}

func LoadVault(vaultPath string) (*Vault, error) {
	var notes []*Note

	err := filepath.WalkDir(vaultPath, func(path string, _ fs.DirEntry, e error) error {
		if e != nil {
			return e
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		note, err := LoadNote(path)
		if err != nil {
			return fmt.Errorf("Failed to read note %s, error %s",
				path, err)
		}

		notes = append(notes, note)
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Id < notes[j].Id
	})

	vault := &Vault{
		Path:  vaultPath,
		Notes: notes,
	}

	// Read from a directory
	return vault, nil
}

func (v *Vault) NextId() int {
	return v.Notes[len(v.Notes)-1].Id + 1
}

func (v *Vault) CreateNote() (*Note, error) {
	newId := v.NextId()

	filename := fmt.Sprintf("%d.md", newId)
	notePath := path.Join(v.Path, filename)

	note, err := CreateNote(notePath, newId)
	if err != nil {
		return nil, err
	}

	v.Notes = append(v.Notes, note)

	err = Edit(notePath)
	if err != nil {
		return nil, err
	}

	return note, nil
}

func (v *Vault) GetNote(id int) *Note {
	for _, note := range v.Notes {
		if note.Id == id {
			return note
		}
	}

	return nil
}
