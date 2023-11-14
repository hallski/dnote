package render

import (
	"dnote/core"
	"fmt"
	"io"
)

// Render a list of links
func LinkList(out io.Writer, lister core.NoteLister, links *core.DocLinks, linkOffset int, styles LinkListStyles) {
	for i, note := range lister.ListNotes() {
		link := links.GetLink(note.ID)
		active := links.IsActive(linkOffset + i)
		fmt.Fprintf(out, "  • %s%s\n",
			RenderLink(link, active, styles.linkStyles),
			styles.titleStyle.Render(" "+note.Title))
	}
}

func RenderLink(link core.ShortcutLink, active bool, styles LinkStyles) string {
	var style = styles.Inactive
	if active {
		style = styles.Active
	}

	if link.Shortcut == "" {
		return styles.Bracket.Render("[[") +
			style.Render(link.ID) +
			styles.Bracket.Render("]]")
	}

	return styles.Bracket.Render("[") +
		styles.Shortcut.Render(link.Shortcut) +
		styles.Bracket.Render("|") +
		style.Render(link.ID) +
		styles.Bracket.Render("]")
}
