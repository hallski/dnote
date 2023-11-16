package render

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/ansi"
)

const TitleBarHeight = 2

func Titlebar(width int, lastId string) string {

	style := lipgloss.NewStyle().Foreground(DividerColor)

	start := style.Render("─/ ") +
		StyleHighRed.Render("Thinkadus") +
		style.Render(" /")

	end := style.Render("/ l.id ") +
		StyleHighCyan.Render(lastId) +
		style.Render(" /─")

	startLen := ansi.PrintableRuneWidth(start)
	endLen := ansi.PrintableRuneWidth(end)
	padLen := max(0, width-startLen-endLen)
	padding := strings.Repeat("─", padLen)

	return start + style.Render(padding) + end + "\n"
}
