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
	Back     key.Binding
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
	Back: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp("b", "back"),
	),
}

type model struct {
	noteBook *mdfiles.MdDirectory
	msg      string

	navStack []string

	width  int
	height int

	doc docModel
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	return model{
		noteBook,
		"Hello there",
		[]string{},
		0, 0,
		newDoc(0, 0),
	}
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
		case key.Matches(msg, DefaultKeyMap.Back):
			m.goBack()
			return m, nil
		}
		var cmd tea.Cmd
		m.doc, cmd = m.doc.Update(msg)
		return m, cmd
	case openLinkMsg:
		m.openNote(msg.id, true)
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

func (m *model) goBack() {
	if len(m.navStack) < 2 {
		fmt.Println("No navstack")
		return
	}

	// TODO: Store the doc model on the stack instead of just an index
	//       this will keep the scroll position as well
	backIdx := len(m.navStack) - 2
	id := m.navStack[backIdx]
	m.navStack = m.navStack[:backIdx+1]
	m.openNote(id, false)
	m.msg = fmt.Sprintf("Stack: %#v", m.navStack)
}

func (m *model) openNote(id string, nav bool) {
	m.msg = "Open" + id
	note := m.noteBook.FindNote(id)
	if note != nil {
		m.doc.renderNote(note)
		if nav {
			m.navStack = append(m.navStack, note.ID)
		}
	}
}

func Run(noteBook *mdfiles.MdDirectory, openId string) error {
	m := initialModel(noteBook)
	m.openNote(openId, true)
	p := tea.NewProgram(m)

	_, err := p.Run()

	return err
}
