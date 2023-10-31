package dnote

type Note struct {
	Id      int
	Path    string
	Title   string
	Content string
}

type NoteStorage interface {
	AllNotes() []*Note
	FindNote(id int) *Note

	CreateNote() (*Note, error)
	DeleteNote(id int) error
}

// Unused
type NoteEditor interface {
	Edit(id int, storage NoteStorage) error
}

type NoteLister interface {
	List(storage NoteStorage) error
}

type NoteViewer interface {
	View(id int, storage NoteStorage) error
}
