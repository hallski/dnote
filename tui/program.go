package tui

import (
	"dnote/mdfiles"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type appKeyMap struct {
	Quit     key.Binding
	Search   key.Binding
	AddNote  key.Binding
	EditNode key.Binding
}

var DefaultKeyMap = appKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Search: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "search"),
	),
	AddNote: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add note"),
	),
	EditNode: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit note"),
	),
}

type model struct {
	noteBook *mdfiles.MdDirectory
	msg      string

	width  int
	height int

	doc docModel
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	return model{
		noteBook,
		"Hello there",
		0, 0,
		newDoc(0, 0, noteBook.FindNote("074"))}
}

func (m model) Init() tea.Cmd {
	// Trigger any initial command
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap.Search):
			m.msg = "Searching!"
			return m, nil
		case key.Matches(msg, DefaultKeyMap.AddNote):
			m.msg = "Add new note"
			return m, nil
		case key.Matches(msg, DefaultKeyMap.EditNode):
			m.msg = "Edit note"
			return m, nil
		}
		var cmd tea.Cmd
		m.doc, cmd = m.doc.Update(msg)
		return m, cmd
	case openLinkMsg:
		m.msg = fmt.Sprintf("Opening %s", msg.id)
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.doc.setSize(m.width, m.height-1)
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	// Render the entire UI
	return lipgloss.JoinVertical(0, m.doc.View(), m.msg)
}

func Run(noteBook *mdfiles.MdDirectory) error {
	p := tea.NewProgram(initialModel(noteBook))

	_, err := p.Run()

	return err
}
