package render

import (
	"dnote/core"
	"fmt"
	"strings"
)

func BackLinks(note *core.Note, docLinkIdx int, docLinks *core.DocLinks, width int) string {

	// Crude backlink support
	builder := new(strings.Builder)
	if len(note.BackLinks.ListNotes()) > 0 {
		bls := new(strings.Builder)

		beforeText := "─ Backlinks "
		beforeLen := width - len([]rune(beforeText))
		if beforeLen > 0 {
			border := strings.Repeat("─", beforeLen)
			fmt.Fprintln(bls, BacklinksTitleStyle.Render(beforeText+border+"\n"))
		}

		LinkList(bls, note.BackLinks, docLinks, docLinkIdx, BackLinkListStyles)

		box := BacklinksBoxStyle.Copy().
			Width(width - BacklinksBoxStyle.GetHorizontalBorderSize())

		fmt.Fprintf(builder, box.Render(bls.String()))
	}

	return builder.String()
}
