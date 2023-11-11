package tui

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

type docLinks struct {
	links []string

	curLink int
}

func newDocLinks(links []string) docLinks {
	return docLinks{
		links,
		0,
	}
}

func (l *docLinks) cycle(dir cycleDirection) {
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

func (l *docLinks) Next() {
	l.cycle(forward)
}

func (l *docLinks) Prev() {
	l.cycle(backward)
}

func (l *docLinks) Current() string {
	if l.curLink >= len(l.links) {
		return ""
	}
	return l.links[l.curLink]
}

func (l *docLinks) GetLinkIdx(idx int) string {
	if idx < 0 || idx >= len(l.links) || idx >= len(shortcuts) {
		return ""
	}

	return l.links[idx]
}

func (l *docLinks) GetShortcut(idx int) string {
	if idx < 0 || idx >= len(l.links) || idx >= len(shortcuts) {
		return ""
	}

	return string(shortcuts[idx])
}

func (l *docLinks) GetLink(shortcut string) string {
	for i, sc := range shortcuts {
		if string(sc) == shortcut {
			if i < len(l.links) {
				return l.links[i]
			}
		}
	}

	return ""
}

func (l *docLinks) IsActive(idx int) bool {
	return idx == l.curLink
}
