package render

import "github.com/charmbracelet/lipgloss"

func BarGraphics(text string) string {
	highlight := lipgloss.NewStyle().Foreground(lipgloss.Color("#4c2ab9"))
	textColor := StyleDivider.Copy().Foreground(ColorWhite).Background(DividerColor)

	return StyleDivider.Render("▀ ") +
		highlight.Render("▀") +
		StyleDivider.Render("▄▀▄██") +
		textColor.Render(text) +
		StyleDivider.Render("██▀▄▀") +
		highlight.Render("▄") +
		StyleDivider.Render(" ▄")
}
