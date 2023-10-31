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

type NoteEditor interface {
	Edit(note *Note) error
}

type NoteLister interface {
	List(note []*Note) error
}
