package render

import (
	"dnote/core"
	"fmt"
	"io"
	"strings"

	"github.com/muesli/ansi"
)

// Render a list of links
func LinkList(
	out io.Writer,
	lister core.NoteLister,
	links *core.DocLinks,
	linkOffset int,
	showTags bool,
	styles LinkListStyles) {
	for i, note := range lister.ListNotes() {
		link := links.GetLink(note.ID)
		active := links.IsActive(linkOffset + i)
		indentPlusLink := fmt.Sprintf("  • %s ",
			RenderLink(link, active, styles.linkStyles))
		fmt.Fprintf(out, "%s%s\n", indentPlusLink, styles.titleStyle.Render(note.Title))
		if showTags && len(note.Tags) > 0 {
			tagsIndent := strings.Repeat(" ", ansi.PrintableRuneWidth(indentPlusLink))
			tags := strings.Join(note.Tags, ", ")
			fmt.Fprintf(out, "%s%s\n", tagsIndent, styles.tagStyle.Render(tags))
		}
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
