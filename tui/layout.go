package tui

type rect struct {
	width  int
	height int
}

type layout struct {
	full rect

	horizontal bool
	showList   bool

	statusBar rect
	listThing rect
	doc       rect
}

const (
	docMinWidth   = 80
	docMinHeight  = 25
	listMaxWidth  = 70
	listMaxHeight = 10
)

func newLayout(width, height int) layout {
	full := rect{width, height}
	statusBar := rect{width, 2}

	if width >= 130 {
		listWidth := width - docMinWidth
		listWidth = min(listWidth, listMaxWidth)

		return layout{
			full,
			true,
			true,
			statusBar,
			rect{listWidth, height - statusBar.height},
			rect{width - listWidth, height - statusBar.height},
		}
	}

	listHeight := height - statusBar.height - docMinHeight
	listHeight = min(listHeight, listMaxHeight)

	return layout{
		full,
		false,
		true,
		statusBar,
		rect{width, listHeight},
		rect{width, height - listHeight - statusBar.height},
	}
}

func (l layout) WithList(list bool) layout {
	if list {
		return newLayout(l.full.width, l.full.height)
	}

	if l.horizontal {
		l.doc.width = l.full.width
	} else {
		l.doc.height += l.listThing.height
	}

	l.showList = false
	return l
}
