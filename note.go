package dnote

type Note struct {
	ID      string
	Path    string
	Title   string
	Content string
	Tags    []string
}

type NoteLister interface {
	ListNotes() []*Note
}

const IDLength = 3
