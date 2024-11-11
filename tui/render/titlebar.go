package render

import (
	"dnote/ext"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/ansi"
)

const TitleBarHeight = 2

func Titlebar(width int, lastId string, gitStatus ext.GitStatus) string {
	gitStatusStyle := GitCleanStyle
	if gitStatus == ext.Dirty {
		gitStatusStyle = GitDirtyStyle
	} else if gitStatus == ext.Updating {
		gitStatusStyle = GitUpdatingStyle
	}

	style := lipgloss.NewStyle().Foreground(TitleBarColor)

	start := BarGraphics("Thinkadus")

	end := style.Render("─[ ") +
		gitStatusStyle.Render("") +
		StyleDarkGray.Render(" :: ") +
		StyleHighCyan.Render(lastId) +
		style.Render(" ]─")

	startLen := ansi.PrintableRuneWidth(start)
	endLen := ansi.PrintableRuneWidth(end)
	padLen := max(0, width-startLen-endLen)
	padding := strings.Repeat("─", padLen)

	return start + style.Render(padding) + end + "\n"
}
