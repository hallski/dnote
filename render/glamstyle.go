package render

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
)

var s = glamour.DarkStyleConfig

func GetGlamming() ansi.StyleConfig {
	s.H1.Prefix = "# "
	s.H1.BackgroundColor = stringPtr("")
	s.H1.Color = stringPtr("#ffff66")

	return s
}

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }
