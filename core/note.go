package core

import (
	"fmt"
	"io"
	"unicode"
)

type Note struct {
	ID        string
	Path      string
	Title     string
	Content   string
	Tags      []string
	Links     []string
	BackLinks []*Note
}

type NoteLister interface {
	ListNotes() []*Note
}

type NoteCreator interface {
	CreateNote(title string) (*Note, error)
}

const IDLength = 3

// https://stackoverflow.com/a/73939904
func EllipticalTruncate(text string, maxLen int) string {
	// Make room for ...
	maxLen -= 3
	lastSpaceIx := maxLen
	len := 0
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		len++
		if len > maxLen {
			return text[:lastSpaceIx] + "..."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}

func ListNoteLinks(lister NoteLister, out io.Writer) {
	const linkLen = 4 + IDLength
	const maxLen = 80

	for _, note := range lister.ListNotes() {
		// To support "- Title......[[ID]]" style link
		truncated := EllipticalTruncate(note.Title, maxLen-linkLen)
		// padLen := maxLen - linkLen - len([]rune(truncated))
		// dots := strings.Repeat(".", padLen)

		fmt.Fprintf(out, "- [[%s]] %s\n", note.ID, truncated)
	}
}
