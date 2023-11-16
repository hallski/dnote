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
	end := StyleDivider.Render("[ ") +
		content +
		StyleDivider.Render(" ]─")

	endLen := ansi.PrintableRuneWidth(end)
	padLen := max(0, width-endLen)
	return "\n" + style.Render(strings.Repeat("─", padLen)) + end
}

func BottomBarNote(note *core.Note, width int) string {
	info := style.Render("id ") +
		CurrentIdStyle.Render(note.ID)
	return renderBar(info, width)
}

func BottomBarSearch(result *search.Result, width int) string {
	hits := strconv.Itoa(len(result.ListNotes()))
	info :=
		style.Render("hits ") +
			NrHitsStyle.Render(hits)
	return renderBar(info, width)
}
