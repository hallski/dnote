package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type searchModel struct {
	width, height int

	query string
}

func newSearchModel() searchModel {
	return searchModel{}
}

func (m searchModel) Init() tea.Cmd {
	return nil
}

func (m searchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m searchModel) View() string {
	style := lipgloss.NewStyle().Width(m.width).Height(m.height).Align(lipgloss.Center)
	return style.Render("Search: " + m.query)
}

func (m *searchModel) setSize(width, height int) {
	m.width, m.height = width, height
}

func (m *searchModel) setQuery(query string) {
	m.query = query
}
