package tui

import (
	"dnote/mdfiles"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type model struct {
	noteBook *mdfiles.MdDirectory
	msg      string

	viewport viewport.Model
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	viewport := viewport.New(80, 25)
	return model{noteBook, "Hello there", viewport}
}

func (m model) Init() tea.Cmd {
	// Trigger any initial command
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	md, err := glamour.Render(m.noteBook.FindNote("001").Content, "dark")
	if err != nil {
		panic(err)
	}
	m.viewport.SetContent(md)
	// Render the entire UI
	return m.viewport.View()
}

func Run(noteBook *mdfiles.MdDirectory) error {
	p := tea.NewProgram(initialModel(noteBook))

	_, err := p.Run()

	return err
}
