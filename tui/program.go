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

	width  int
	height int

	viewport viewport.Model
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	viewport := viewport.New(0, 0)
	return model{noteBook, "Hello there", 0, 0, viewport}
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
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		return m.SetupViewport(), nil
	}

	return m, nil
}

func (m model) SetupViewport() tea.Model {
	m.viewport = viewport.New(m.width, m.height)

	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(m.width),
	)
	md, err := r.Render(m.noteBook.FindNote("074").Content)
	if err != nil {
		panic(err)
	}
	m.viewport.SetContent(md)

	return m
}

func (m model) View() string {
	// Render the entire UI
	return m.viewport.View()
}

func Run(noteBook *mdfiles.MdDirectory) error {
	p := tea.NewProgram(initialModel(noteBook))

	_, err := p.Run()

	return err
}
