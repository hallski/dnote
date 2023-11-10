package tui

import (
	"dnote/ext"
	"dnote/mdfiles"
	"fmt"
	"os/exec"

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
	Forward  key.Binding
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
		key.WithKeys("b", "["),
		key.WithHelp("b", "back"),
	),
	Forward: key.NewBinding(
		key.WithKeys("]", "f"),
		key.WithHelp("ctrl+i or f", "forward"),
	),
}

type model struct {
	noteBook *mdfiles.MdDirectory

	msg string

	history *history[string]

	width  int
	height int

	doc docModel
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	return model{
		noteBook,
		"Hello there",
		NewHistory[string](),
		0, 0,
		newDoc(0, 0),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

type statusMsg struct{ s string }
type refreshNotebook struct{}

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
			return m, openEditor(m.noteBook, m.history.GetCurrent())
		case key.Matches(msg, DefaultKeyMap.Back):
			id := m.history.GoBack()
			m.openNote(id, false)
			return m, nil
		case key.Matches(msg, DefaultKeyMap.Forward):
			id := m.history.GoForward()
			m.openNote(id, false)
			return m, nil
		}
		var cmd tea.Cmd
		m.doc, cmd = m.doc.Update(msg)
		return m, cmd
	case statusMsg:
		m.msg = msg.s
		return m, nil
	case refreshNotebook:
		m.refreshNotebook()
		return m, nil
	case openLinkMsg:
		m.openNote(msg.id, true)
		//		m.msg = fmt.Sprintf("Opening %s", msg.id)
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

func openEditor(noteBook *mdfiles.MdDirectory, id string) tea.Cmd {
	note := noteBook.FindNote(id)
	if note == nil {
		return func() tea.Msg { return statusMsg{"Failed opening " + id} }
	}

	editor := ext.GetEditor()
	c := exec.Command(editor, note.Path)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			return statusMsg{fmt.Sprintf("Failed editing: %s", err)}
		} else {
			return refreshNotebook{}
		}
	})
}

func (m *model) refreshNotebook() {
	noteBook, err := mdfiles.Load(m.noteBook.Path)
	if err != nil {
		panic(err)
	}

	m.noteBook = noteBook
	// Force a rerender of the document
	if m.history.GetCurrent() != "" {
		note := noteBook.FindNote(m.history.GetCurrent())
		m.doc.renderNote(note)
	}
}

func (m *model) openNote(id string, nav bool) {
	note := m.noteBook.FindNote(id)
	if note != nil {
		m.doc.renderNote(note)
		if nav {
			m.history.Push(id)
		}
	}
	m.msg = fmt.Sprintf("HISTORY: %v", m.history.stack)
}

func Run(noteBook *mdfiles.MdDirectory, openId string) error {
	m := initialModel(noteBook)
	m.openNote(openId, true)
	p := tea.NewProgram(m)

	_, err := p.Run()

	return err
}
