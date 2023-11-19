package render

import (
	"dnote/core"
	"dnote/search"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/ansi"
)

var style = lipgloss.NewStyle().Foreground(ColorDarkGray)

const BottomBarHeight = 2

func renderBar(content string, width int) string {
	start := StyleDivider.Render("─[ ") +
		content +
		StyleDivider.Render(" ]")

	startLen := ansi.PrintableRuneWidth(start)
	padLen := max(0, width-startLen)
	return "\n" + start + style.Render(strings.Repeat("─", padLen))
}

func BottomBarNote(note *core.Note, width int) string {
	info := CurrentIdStyle.Render(note.ID) +
		StyleDivider.Render("::") +
		CurrentDateStyle.Render(note.Date.Format("2006-01-02"))

	return renderBar(info, width)
}

func BottomBarSearch(result *search.Result, width int) string {
	hits := strconv.Itoa(len(result.ListNotes()))
	info :=
		NrHitsStyle.Render(hits) +
			style.Render(" hits")

	return renderBar(info, width)
}
