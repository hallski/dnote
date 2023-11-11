package tui

import "github.com/charmbracelet/lipgloss"

var backlinksBackgroundStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#222222"))

var backlinksTitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#aaaaaa")).
	MarginBottom(1).
	Underline(true).
	Inherit(backlinksBackgroundStyle)

var backlinksBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true, false).
	BorderBackground(lipgloss.Color("#222222")).
	BorderForeground(lipgloss.Color("#555555")).
	Padding(0, 2, 0, 2).
	Inherit(backlinksBackgroundStyle)

type linkStyles struct {
	inactive lipgloss.Style
	active   lipgloss.Style
	shortcut lipgloss.Style
	bracket  lipgloss.Style
}

var docLinkStyles = linkStyles{
	inactive: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#aa00aa")),
	active: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffff55")).
		Bold(true),
	shortcut: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#55ff55")),
	bracket: lipgloss.NewStyle().
		Foreground(lipgloss.Color("#555555")),
}

var backLinkStyles = linkStyles{
	inactive: docLinkStyles.inactive.Inherit(backlinksBackgroundStyle),
	active:   docLinkStyles.active.Inherit(backlinksBackgroundStyle),
	shortcut: docLinkStyles.shortcut.Inherit(backlinksBackgroundStyle),
	bracket:  docLinkStyles.bracket.Inherit(backlinksBackgroundStyle),
}
