package mdfiles

import (
	"dnote"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"sort"
)

type MdDirectory struct {
	Path  string
	notes []*dnote.Note
}

func noteLoader(notes *[]*dnote.Note) fs.WalkDirFunc {
	return func(path string, _ fs.DirEntry, e error) error {
		if e != nil {
			return e
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		note, err := loadNote(path)
		if err != nil {
			return fmt.Errorf("Failed to read note %s, error %s", path, err)
		}

		*notes = append(*notes, note)
		return nil
	}
}

func Load(dir string) (*MdDirectory, error) {
	var notes []*dnote.Note

	err := filepath.WalkDir(dir, noteLoader(&notes))
	if err != nil {
		return nil, err
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Id < notes[j].Id
	})

	mdd := &MdDirectory{
		Path:  dir,
		notes: notes,
	}

	return mdd, nil
}

func (mdd *MdDirectory) nextId() int {
	return mdd.notes[len(mdd.notes)-1].Id + 1
}

func (mdd *MdDirectory) CreateNote() (*dnote.Note, error) {
	id := mdd.nextId()

	filename := fmt.Sprintf("%d.md", id)
	path := path.Join(mdd.Path, filename)

	note, err := createNote(path, id)
	if err != nil {
		return nil, err
	}

	mdd.notes = append(mdd.notes, note)

	return note, nil
}

func (mdd *MdDirectory) FindNote(id int) *dnote.Note {
	for _, note := range mdd.notes {
		if note.Id == id {
			return note
		}
	}

	return nil
}

func (mdd *MdDirectory) ListNotes() []*dnote.Note {
	return mdd.notes
}

func (mdd *MdDirectory) DeleteNote(id int) error {
	// TODO: Implement
	return nil
}
