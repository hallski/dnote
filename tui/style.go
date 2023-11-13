package tui

import "github.com/charmbracelet/lipgloss"

var (
	colorBlack       = lipgloss.Color("000000")
	colorLowRed      = lipgloss.Color("#aa0000")
	colorLowGreen    = lipgloss.Color("#00aa00")
	colorBrown       = lipgloss.Color("#aa5500")
	colorLowBlue     = lipgloss.Color("#0000aa")
	colorLowMagenta  = lipgloss.Color("#aa00aa")
	colorLowCyan     = lipgloss.Color("#00aaaa")
	colorLightGray   = lipgloss.Color("#aaaaaa")
	colorDarkGray    = lipgloss.Color("#555555")
	colorHighRed     = lipgloss.Color("#ff5555")
	colorHighGreen   = lipgloss.Color("#55ff55")
	colorYellow      = lipgloss.Color("#ffff55")
	colorHighBlue    = lipgloss.Color("#5555ff")
	colorHighCyan    = lipgloss.Color("#55ffff")
	colorHighMagenta = lipgloss.Color("#ff55ff")
	colorWhite       = lipgloss.Color("#ffffff")
)

var colorPanelBackground = lipgloss.Color("#222222")

var backlinksBackgroundStyle = lipgloss.NewStyle().
	Background(colorPanelBackground)

var backlinksLinkTitlestyle = lipgloss.NewStyle().
	Foreground(colorWhite).
	Background(lipgloss.Color("#222222"))

var backlinksTitleStyle = lipgloss.NewStyle().
	Foreground(colorLightGray).
	MarginBottom(1).
	Underline(true).
	Inherit(backlinksBackgroundStyle)

var backlinksBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true, false).
	BorderBackground(colorPanelBackground).
	BorderForeground(lipgloss.Color(colorDarkGray)).
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
		Foreground(colorLowGreen),
	active: lipgloss.NewStyle().
		Foreground(colorYellow),
	shortcut: lipgloss.NewStyle().
		Foreground(colorHighGreen),
	bracket: lipgloss.NewStyle().
		Foreground(colorDarkGray),
}

var backLinkStyles = linkStyles{
	inactive: docLinkStyles.inactive.Copy().Inherit(backlinksBackgroundStyle),
	active:   docLinkStyles.active.Copy().Inherit(backlinksBackgroundStyle),
	shortcut: docLinkStyles.shortcut.Copy().Inherit(backlinksBackgroundStyle),
	bracket:  docLinkStyles.bracket.Copy().Inherit(backlinksBackgroundStyle),
}

var docNoteIdStyle = lipgloss.NewStyle().
	Foreground(colorHighGreen)

var currentIdStyle = lipgloss.NewStyle().Foreground(colorYellow)
