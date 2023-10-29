package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LoadingScreenModel struct {
	spinner spinner.Model
}

func (m LoadingScreenModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m LoadingScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m LoadingScreenModel) View() string {
	return fmt.Sprintf("\n\n   %s Loading notes...\n\n", m.spinner.View())
}

func InitLoadingScreen() tea.Model {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return LoadingScreenModel{spinner: s}
}
