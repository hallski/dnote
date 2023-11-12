package mdfiles

import (
	"fmt"
	"io/fs"
	"math/rand"
	"path"
	"path/filepath"
	"sort"
	"strconv"

	"dnote/core"
)

type BackLinks map[string][]*core.Note

type MdDirectory struct {
	Path  string
	notes []*core.Note
}

func noteLoader(notes *[]*core.Note) fs.WalkDirFunc {
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
	var notes []*core.Note

	err := filepath.WalkDir(dir, noteLoader(&notes))
	if err != nil {
		return nil, err
	}

	backlinks := make(BackLinks)

	for _, note := range notes {
		for _, link := range note.Links {
			backlinks[link] = append(backlinks[link], note)
		}
	}

	for _, note := range notes {
		note.BackLinks = backlinks[note.ID]
	}

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].ID < notes[j].ID
	})

	mdd := &MdDirectory{
		Path:  dir,
		notes: notes,
	}

	return mdd, nil
}

func (mdd *MdDirectory) nextID() string {
	idx := len(mdd.notes) - 1
	lastID, err := strconv.Atoi(mdd.notes[idx].ID)
	if err != nil {
		return "err"
	}
	return PadID(strconv.Itoa(lastID + 1))
}

func (mdd *MdDirectory) notePath(id string) string {
	filename := fmt.Sprintf("%s.md", id)
	return path.Join(mdd.Path, filename)
}

func (mdd *MdDirectory) CreateNote(title string) (*core.Note, error) {
	id := mdd.nextID()

	filename := fmt.Sprintf("%s.md", id)
	path := path.Join(mdd.Path, filename)

	note, err := createNote(path, id, title)
	if err != nil {
		return nil, err
	}

	mdd.notes = append(mdd.notes, note)

	return note, nil
}

func (mdd *MdDirectory) FindNote(id string) *core.Note {
	if id == "last" {
		return mdd.LastNote()
	}
	id = PadID(id)
	for _, note := range mdd.notes {
		if note.ID == id {
			return note
		}
	}

	return nil
}

func (mdd *MdDirectory) ListNotes() []*core.Note {
	return mdd.notes
}

func (mdd *MdDirectory) DeleteNote(id int) error {
	// TODO: Implement
	return nil
}

func (mdd *MdDirectory) RandomNote() *core.Note {
	return mdd.notes[rand.Intn(len(mdd.notes))]
}

func (mdd *MdDirectory) LastNote() *core.Note {
	return mdd.notes[len(mdd.notes)-1]
}

func (mdd *MdDirectory) Rename(oldID, newID string) error {
	return changeID(mdd, oldID, newID)
}

func (mdd *MdDirectory) Migrate() error {
	for _, note := range mdd.notes {
		newID := PadID(note.ID)
		if newID != note.ID {
			mdd.Rename(note.ID, newID)
		}
	}

	return nil
}

type Result struct {
	result []*core.Note
}

func (sr *Result) ListNotes() []*core.Note {
	return sr.result
}

// Should this be in search, or just a generic TurnListOfIdsIntoListOfNotes?
func (mdd *MdDirectory) Backlinks(id string) *Result {
	note := mdd.FindNote(id)

	return &Result{note.BackLinks}
}

func (mdd *MdDirectory) Orphans() *Result {
	var orphans []*core.Note
	for _, note := range mdd.notes {
		if len(note.BackLinks) == 0 {
			orphans = append(orphans, note)
		}
	}

	return &Result{orphans}
}
