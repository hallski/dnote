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
	s.H1.Color = stringPtr(cfg.Theme.Glamour.H1Color)

	s.BlockQuote.Italic = boolPtr(true)

	return s
}

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }
