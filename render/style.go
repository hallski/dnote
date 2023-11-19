package render

import "github.com/charmbracelet/lipgloss"

var (
	ColorBlack       = lipgloss.Color("#000000")
	ColorLowRed      = lipgloss.Color("#aa0000")
	ColorLowGreen    = lipgloss.Color("#00aa00")
	ColorBrown       = lipgloss.Color("#aa5500")
	ColorLowBlue     = lipgloss.Color("#0000aa")
	ColorLowMagenta  = lipgloss.Color("#aa00aa")
	ColorLowCyan     = lipgloss.Color("#00aaaa")
	ColorLightGray   = lipgloss.Color("#aaaaaa")
	ColorDarkGray    = lipgloss.Color("#555555")
	ColorHighRed     = lipgloss.Color("#ff5555")
	ColorHighGreen   = lipgloss.Color("#55ff55")
	ColorYellow      = lipgloss.Color("#ffff55")
	ColorHighBlue    = lipgloss.Color("#5555ff")
	ColorHighCyan    = lipgloss.Color("#55ffff")
	ColorHighMagenta = lipgloss.Color("#ff55ff")
	ColorWhite       = lipgloss.Color("#ffffff")
)

var StyleBlack = lipgloss.NewStyle().Foreground(ColorBlack)
var StyleLowRed = lipgloss.NewStyle().Foreground(ColorLowRed)
var StyleLowGreen = lipgloss.NewStyle().Foreground(ColorLowGreen)
var StyleBrown = lipgloss.NewStyle().Foreground(ColorBrown)
var StyleLowBlue = lipgloss.NewStyle().Foreground(ColorLowBlue)
var StyleLowMagenta = lipgloss.NewStyle().Foreground(ColorLowMagenta)
var StyleLowCyan = lipgloss.NewStyle().Foreground(ColorLowCyan)
var StyleLightGray = lipgloss.NewStyle().Foreground(ColorLightGray)
var StyleDarkGray = lipgloss.NewStyle().Foreground(ColorDarkGray)
var StyleHighRed = lipgloss.NewStyle().Foreground(ColorHighRed)
var StyleHighGreen = lipgloss.NewStyle().Foreground(ColorHighGreen)
var StyleYellow = lipgloss.NewStyle().Foreground(ColorYellow)
var StyleHighBlue = lipgloss.NewStyle().Foreground(ColorHighBlue)
var StyleHighCyan = lipgloss.NewStyle().Foreground(ColorHighCyan)
var StyleHighMagenta = lipgloss.NewStyle().Foreground(ColorHighMagenta)
var StyleWhite = lipgloss.NewStyle().Foreground(ColorWhite)

var DividerColor = ColorDarkGray

var StyleDivider = lipgloss.NewStyle().Foreground(DividerColor)
var ColorPanelBackground = lipgloss.Color("#222222")

var BacklinksBackgroundStyle = lipgloss.NewStyle().
	Background(ColorPanelBackground).
	Foreground(ColorDarkGray)

var BacklinksLinkTitlestyle = lipgloss.NewStyle().
	Foreground(ColorWhite).
	Background(lipgloss.Color("#222222"))

var BacklinksTitleStyle = lipgloss.NewStyle().
	Foreground(ColorDarkGray).
	Inherit(BacklinksBackgroundStyle)

var BacklinksBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false).
	BorderBottom(true).
	BorderBackground(ColorPanelBackground).
	BorderForeground(ColorDarkGray).
	Inherit(BacklinksBackgroundStyle)

type LinkStyles struct {
	Inactive lipgloss.Style
	Active   lipgloss.Style
	Shortcut lipgloss.Style
	Bracket  lipgloss.Style
}

type LinkListStyles struct {
	linkStyles LinkStyles
	titleStyle lipgloss.Style
}

var DocLinkStyles = LinkStyles{
	Inactive: lipgloss.NewStyle().
		Foreground(ColorLowGreen),
	Active: lipgloss.NewStyle().
		Foreground(ColorYellow),
	Shortcut: lipgloss.NewStyle().
		Foreground(ColorHighGreen),
	Bracket: lipgloss.NewStyle().
		Foreground(ColorDarkGray),
}

var BackLinkStyles = LinkStyles{
	Inactive: DocLinkStyles.Inactive.Copy().Inherit(BacklinksBackgroundStyle),
	Active:   DocLinkStyles.Active.Copy().Inherit(BacklinksBackgroundStyle),
	Shortcut: DocLinkStyles.Shortcut.Copy().Inherit(BacklinksBackgroundStyle),
	Bracket:  DocLinkStyles.Bracket.Copy().Inherit(BacklinksBackgroundStyle),
}

var DocLinkListStyles = LinkListStyles{
	DocLinkStyles,
	lipgloss.NewStyle().Foreground(ColorLightGray),
}

var BackLinkListStyles = LinkListStyles{
	BackLinkStyles,
	BacklinksLinkTitlestyle,
}

var DocNoteIdStyle = lipgloss.NewStyle().
	Foreground(ColorHighGreen)

var CurrentDateStyle = lipgloss.NewStyle().Foreground(ColorLowCyan)
var CurrentIdStyle = lipgloss.NewStyle().Foreground(ColorYellow)
var NrHitsStyle = lipgloss.NewStyle().Foreground(ColorYellow)

var TagsStyle = lipgloss.NewStyle().Foreground(ColorHighBlue)

var GitCleanStyle = StyleDarkGray.Copy()
var GitDirtyStyle = StyleYellow.Copy()
var GitUpdatingStyle = StyleHighMagenta.Copy()

var TitleBarColor = lipgloss.Color("#391f8b")
var TitleBarTextColor = ColorWhite
