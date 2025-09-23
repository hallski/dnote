package render

import (
	"dnote/config"

	"github.com/charmbracelet/lipgloss"
)

// Color variables - these will be initialized from config
var (
	ColorBlack           lipgloss.Color
	ColorLowRed          lipgloss.Color
	ColorLowGreen        lipgloss.Color
	ColorBrown           lipgloss.Color
	ColorLowBlue         lipgloss.Color
	ColorLowMagenta      lipgloss.Color
	ColorLowCyan         lipgloss.Color
	ColorLightGray       lipgloss.Color
	ColorDarkGray        lipgloss.Color
	ColorHighRed         lipgloss.Color
	ColorHighGreen       lipgloss.Color
	ColorYellow          lipgloss.Color
	ColorHighBlue        lipgloss.Color
	ColorHighCyan        lipgloss.Color
	ColorHighMagenta     lipgloss.Color
	ColorWhite           lipgloss.Color
	DividerColor         lipgloss.Color
	ColorPanelBackground lipgloss.Color
	TitleBarColor        lipgloss.Color
	TitleBarTextColor    lipgloss.Color
	TagsColor            lipgloss.Color
)

// Style variables - these will be initialized after colors
var (
	StyleBlack       lipgloss.Style
	StyleLowRed      lipgloss.Style
	StyleLowGreen    lipgloss.Style
	StyleBrown       lipgloss.Style
	StyleLowBlue     lipgloss.Style
	StyleLowMagenta  lipgloss.Style
	StyleLowCyan     lipgloss.Style
	StyleLightGray   lipgloss.Style
	StyleDarkGray    lipgloss.Style
	StyleHighRed     lipgloss.Style
	StyleHighGreen   lipgloss.Style
	StyleYellow      lipgloss.Style
	StyleHighBlue    lipgloss.Style
	StyleHighCyan    lipgloss.Style
	StyleHighMagenta lipgloss.Style
	StyleWhite       lipgloss.Style
	StyleDivider     lipgloss.Style
)

// Complex styles
var (
	BacklinksBackgroundStyle lipgloss.Style
	BacklinksLinkTitlestyle  lipgloss.Style
	BacklinksTitleStyle      lipgloss.Style
	BacklinksBoxStyle        lipgloss.Style
	DocLinkStyles            LinkStyles
	BackLinkStyles           LinkStyles
	DocLinkListStyles        LinkListStyles
	BackLinkListStyles       LinkListStyles
	DocNoteIdStyle           lipgloss.Style
	CurrentDateStyle         lipgloss.Style
	CurrentIdStyle           lipgloss.Style
	NrHitsStyle              lipgloss.Style
	TagsStyle                lipgloss.Style
	BackLinkCountStyle       lipgloss.Style
	GitCleanStyle            lipgloss.Style
	GitDirtyStyle            lipgloss.Style
	GitUpdatingStyle         lipgloss.Style
)

type LinkStyles struct {
	Inactive lipgloss.Style
	Active   lipgloss.Style
	Shortcut lipgloss.Style
	Bracket  lipgloss.Style
}

type LinkListStyles struct {
	linkStyles         LinkStyles
	titleStyle         lipgloss.Style
	tagStyle           lipgloss.Style
	backLinkCountStyle lipgloss.Style
}

// InitializeStyles initializes all styles from the configuration
// This should be called after the configuration is loaded
func InitializeStyles() {
	cfg := config.GetConfig()
	colors := cfg.Theme.Colors

	// Initialize colors
	ColorBlack = lipgloss.Color(colors.Black)
	ColorLowRed = lipgloss.Color(colors.LowRed)
	ColorLowGreen = lipgloss.Color(colors.LowGreen)
	ColorBrown = lipgloss.Color(colors.Brown)
	ColorLowBlue = lipgloss.Color(colors.LowBlue)
	ColorLowMagenta = lipgloss.Color(colors.LowMagenta)
	ColorLowCyan = lipgloss.Color(colors.LowCyan)
	ColorLightGray = lipgloss.Color(colors.LightGray)
	ColorDarkGray = lipgloss.Color(colors.DarkGray)
	ColorHighRed = lipgloss.Color(colors.HighRed)
	ColorHighGreen = lipgloss.Color(colors.HighGreen)
	ColorYellow = lipgloss.Color(colors.Yellow)
	ColorHighBlue = lipgloss.Color(colors.HighBlue)
	ColorHighCyan = lipgloss.Color(colors.HighCyan)
	ColorHighMagenta = lipgloss.Color(colors.HighMagenta)
	ColorWhite = lipgloss.Color(colors.White)
	DividerColor = lipgloss.Color(colors.Divider)
	ColorPanelBackground = lipgloss.Color(colors.PanelBackground)
	TitleBarColor = lipgloss.Color(colors.TitleBar)
	TitleBarTextColor = lipgloss.Color(colors.TitleBarText)
	TagsColor = lipgloss.Color(colors.Tags)

	// Initialize basic styles
	StyleBlack = lipgloss.NewStyle().Foreground(ColorBlack)
	StyleLowRed = lipgloss.NewStyle().Foreground(ColorLowRed)
	StyleLowGreen = lipgloss.NewStyle().Foreground(ColorLowGreen)
	StyleBrown = lipgloss.NewStyle().Foreground(ColorBrown)
	StyleLowBlue = lipgloss.NewStyle().Foreground(ColorLowBlue)
	StyleLowMagenta = lipgloss.NewStyle().Foreground(ColorLowMagenta)
	StyleLowCyan = lipgloss.NewStyle().Foreground(ColorLowCyan)
	StyleLightGray = lipgloss.NewStyle().Foreground(ColorLightGray)
	StyleDarkGray = lipgloss.NewStyle().Foreground(ColorDarkGray)
	StyleHighRed = lipgloss.NewStyle().Foreground(ColorHighRed)
	StyleHighGreen = lipgloss.NewStyle().Foreground(ColorHighGreen)
	StyleYellow = lipgloss.NewStyle().Foreground(ColorYellow)
	StyleHighBlue = lipgloss.NewStyle().Foreground(ColorHighBlue)
	StyleHighCyan = lipgloss.NewStyle().Foreground(ColorHighCyan)
	StyleHighMagenta = lipgloss.NewStyle().Foreground(ColorHighMagenta)
	StyleWhite = lipgloss.NewStyle().Foreground(ColorWhite)
	StyleDivider = lipgloss.NewStyle().Foreground(DividerColor)

	// Initialize complex styles
	initializeComplexStyles()
}

func initializeComplexStyles() {
	BacklinksBackgroundStyle = lipgloss.NewStyle().
		Background(ColorPanelBackground).
		Foreground(ColorDarkGray)

	BacklinksLinkTitlestyle = lipgloss.NewStyle().
		Foreground(ColorWhite).
		Background(ColorPanelBackground)

	BacklinksTitleStyle = lipgloss.NewStyle().
		Foreground(ColorDarkGray).
		Inherit(BacklinksBackgroundStyle)

	BacklinksBoxStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false).
		BorderBottom(true).
		BorderBackground(ColorPanelBackground).
		BorderForeground(ColorDarkGray).
		Inherit(BacklinksBackgroundStyle)

	DocLinkStyles = LinkStyles{
		Inactive: lipgloss.NewStyle().Foreground(ColorLowGreen),
		Active:   lipgloss.NewStyle().Foreground(ColorYellow),
		Shortcut: lipgloss.NewStyle().Foreground(ColorHighGreen),
		Bracket:  lipgloss.NewStyle().Foreground(ColorDarkGray),
	}

	BackLinkStyles = LinkStyles{
		Inactive: DocLinkStyles.Inactive.Copy().Inherit(BacklinksBackgroundStyle),
		Active:   DocLinkStyles.Active.Copy().Inherit(BacklinksBackgroundStyle),
		Shortcut: DocLinkStyles.Shortcut.Copy().Inherit(BacklinksBackgroundStyle),
		Bracket:  DocLinkStyles.Bracket.Copy().Inherit(BacklinksBackgroundStyle),
	}

	TagsStyle = lipgloss.NewStyle().Foreground(TagsColor)
	BackLinkCountStyle = lipgloss.NewStyle().Foreground(ColorHighRed)

	DocLinkListStyles = LinkListStyles{
		DocLinkStyles,
		lipgloss.NewStyle().Foreground(ColorLightGray),
		TagsStyle,
		BackLinkCountStyle,
	}

	BackLinkListStyles = LinkListStyles{
		BackLinkStyles,
		BacklinksLinkTitlestyle,
		TagsStyle.Copy().Inherit(BacklinksBackgroundStyle),
		BackLinkCountStyle.Copy().Inherit(BacklinksBackgroundStyle),
	}

	DocNoteIdStyle = lipgloss.NewStyle().Foreground(ColorHighGreen)
	CurrentDateStyle = lipgloss.NewStyle().Foreground(ColorLowCyan)
	CurrentIdStyle = lipgloss.NewStyle().Foreground(ColorYellow)
	NrHitsStyle = lipgloss.NewStyle().Foreground(ColorYellow)
	GitCleanStyle = StyleDarkGray.Copy()
	GitDirtyStyle = StyleYellow.Copy()
	GitUpdatingStyle = StyleHighMagenta.Copy()
}
