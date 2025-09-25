package render

import (
	"dnote/config"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
)

func GetGlamming() ansi.StyleConfig {
	cfg := config.GetConfig()
	s := glamour.DarkStyleConfig

	s.H1.Prefix = "# "
	s.H1.BackgroundColor = stringPtr("")
	s.H1.Color = stringPtr(cfg.Theme.Colors.H1Color)
	if cfg.Theme.Colors.H2Color != "" {
		s.H2.Color = stringPtr(cfg.Theme.Colors.H2Color)
	}
	if cfg.Theme.Colors.H3Color != "" {
		s.H3.Color = stringPtr(cfg.Theme.Colors.H3Color)
	}

	s.BlockQuote.Italic = boolPtr(true)

	return s
}

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }
