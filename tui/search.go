package tui

import (
	"dnote/core"
	"dnote/search"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type searchModel struct {
	width, height int

	query  string
	result *search.Result
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
	builder := new(strings.Builder)

	style := lipgloss.NewStyle().Width(m.width).Height(m.height - len(m.result.Result)).Align(lipgloss.Center)
	fmt.Fprint(builder, style.Render("Search: "+m.query))
	core.ListNoteLinks(m.result, builder)

	return builder.String()
}

func (m *searchModel) setSize(width, height int) {
	m.width, m.height = width, height
}

func (m *searchModel) setQuery(query string) {
	m.query = query
}

func (m *searchModel) setResult(result *search.Result) {
	m.result = result
}
