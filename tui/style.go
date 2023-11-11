package tui

import "github.com/charmbracelet/lipgloss"

var linkBracketStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
var linkInactiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#aa00aa"))
var linkActiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff55")).Bold(true)
var linkShortcutStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#55ff55"))
