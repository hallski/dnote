package dnote

type Note struct {
	Id      int
	Path    string
	Title   string
	Content string
	Tags    []string
}

type NoteLister interface {
	ListNotes() []*Note
}
