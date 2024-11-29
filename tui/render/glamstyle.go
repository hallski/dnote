package render

import (
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/glamour/styles"
)

var s = styles.DarkStyleConfig

func GetGlamming() ansi.StyleConfig {
	s.H1.Prefix = "# "
	s.H1.BackgroundColor = stringPtr("")
	s.H1.Color = stringPtr("#ffff66")

	s.BlockQuote.Italic = boolPtr(true)

	return s
}

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
