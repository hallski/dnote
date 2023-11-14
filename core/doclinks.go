package core

// Handling of links in a document
// Nows which one is active and cycles through them
// Assigns and looks up shortcuts

// Empty strings indicate no link or shortcut

var shortcuts = []byte("ABCDEFGHIJKLMNOPQRSTUVXYZ")

type cycleDirection uint

const (
	forward cycleDirection = iota
	backward
)

type ShortcutLink struct {
	ID       string
	Shortcut string
}

type DocLinks struct {
	links []ShortcutLink

	curLink int
}

func NewDocLinks(links []string) DocLinks {
	var ll []ShortcutLink
	m := make(map[string]string)

	first := -1
	i := 0
	for _, linkID := range links {
		sc := ""
		if m[linkID] != "" {
			sc = m[linkID]
		} else if i < len(shortcuts) {
			sc = string(shortcuts[i])
			m[linkID] = sc
			i++
		} else {
			// Position current link on first link without shortcut.
			// This makes it easy to navigate as I can use next to move
			// on from there.
			if first < 0 {
				first = i
			}
		}
		ll = append(ll, ShortcutLink{linkID, sc})
	}
	return DocLinks{ll, first}
}

func (l *DocLinks) cycle(dir cycleDirection) {
	length := len(l.links)

	if length <= 0 {
		return
	}

	if dir == forward {
		l.curLink = (l.curLink + 1) % length
	} else {
		if l.curLink < 0 {
			l.curLink = 0
		}
		l.curLink = (l.curLink + length - 1) % length
	}
}

func (l *DocLinks) Next() {
	l.cycle(forward)
}

func (l *DocLinks) Prev() {
	l.cycle(backward)
}

func (l *DocLinks) Current() string {
	if l.curLink < 0 || l.curLink >= len(l.links) {
		return ""
	}
	return l.links[l.curLink].ID
}

func (l *DocLinks) GetLinkIdx(idx int) ShortcutLink {
	if idx < 0 || idx >= len(l.links) || idx >= len(shortcuts) {
		return ShortcutLink{}
	}

	return l.links[idx]
}

func (l *DocLinks) GetShortcut(id string) string {
	for _, link := range l.links {
		if link.ID == id {
			return link.Shortcut
		}
	}

	return ""
}

func (l *DocLinks) GetShortcutIdx(idx int) string {
	if idx < 0 || idx >= len(l.links) || idx >= len(shortcuts) {
		return ""
	}

	return l.links[idx].Shortcut
}

func (l *DocLinks) GetLink(id string) ShortcutLink {
	for _, link := range l.links {
		if link.ID == id {
			return link
		}
	}

	return ShortcutLink{}
}

func (l *DocLinks) GetLinkFromShortcut(shortcut string) ShortcutLink {
	for _, link := range l.links {
		if link.Shortcut == shortcut {
			return link
		}
	}

	return ShortcutLink{}
}

func (l *DocLinks) IsActive(idx int) bool {
	return idx == l.curLink
}
