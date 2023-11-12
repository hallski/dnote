package tui

import (
	"dnote/core"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type listThing struct {
	size  rect
	notes core.NoteLister
}

func (m *listThing) Init() tea.Cmd { return nil }

func (m *listThing) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *listThing) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorLowBlue).
		Foreground(colorHighBlue).
		Padding(0, 1, 1)

	active := lipgloss.NewStyle().Foreground(colorHighRed)

	return style.
		Width(m.size.width - style.GetBorderLeftSize() - style.GetBorderRightSize()).
		Height(m.size.height - style.GetBorderTopSize() - style.GetBorderBottomSize()).
		Render("List\n" + active.Render("[1] A nice title") + "\n[2] Another note\n[3] This is a third")
}

func (m *listThing) setSize(size rect) {
	m.size = size
}

func newListThing(notes core.NoteLister) listThing {
	return listThing{
		rect{0, 0},
		notes,
	}
}
