package tui

import (
	"dnote/core"
	"dnote/mdfiles"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	statusBarHeight = 2
)

func getStrings(bytes []byte) []string {
	var ss []string

	for _, b := range bytes {
		ss = append(ss, string(b))
	}
	return ss
}

type model struct {
	noteBook *mdfiles.MdDirectory

	keymap appKeyMap

	statusMsg string

	history *history[string]

	width  int
	height int

	doc docModel

	enteringCmd bool
	commandBar  commandBar
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	return model{
		noteBook,
		defaultAppKeyMap,
		"",
		NewHistory[string](),
		0, 0,
		newDoc(0, 0),
		false,
		newCommandBar(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		case m.enteringCmd:
			var cmd tea.Cmd
			m.commandBar, cmd = m.commandBar.Update(msg)
			return m, cmd
		case key.Matches(msg, m.keymap.Search):
			m.statusMsg = "Searching!"
			return m, nil
		case key.Matches(msg, m.keymap.EditNode):
			return m, openEditor(m.noteBook, m.history.GetCurrent())
		case key.Matches(msg, m.keymap.Back):
			id := m.history.GoBack()
			m.openNote(id, false)
			return m, nil
		case key.Matches(msg, m.keymap.Forward):
			id := m.history.GoForward()
			m.openNote(id, false)
			return m, nil
		case key.Matches(msg, m.keymap.StartCmd):
			m.enteringCmd = true
			return m, nil
		case key.Matches(msg, m.keymap.QuickOpen):
			m.commandBar.startOpen(msg.String())
			m.enteringCmd = true
			return m, nil
		}
		var cmd tea.Cmd
		m.doc, cmd = m.doc.Update(msg)
		return m, cmd

	case statusMsg:
		m.statusMsg = msg.s
		return m, nil
	case exitCmdMsg:
		m.enteringCmd = false
		return m, nil
	case editorFinishedMsg:
		return m, refreshNotebook(m.noteBook.Path)
	case refreshNotebookMsg:
		return m, refreshNotebook(m.noteBook.Path)
	case openRandomMsg:
		note := m.noteBook.RandomNote()
		return m, openLinkCmd(note.ID)
	case addNoteMessage:
		note, err := m.noteBook.CreateNote(msg.title)
		if err != nil {
			m.statusMsg = "Error creating note"
			return m, nil
		}
		return m, openEditor(m.noteBook, note.ID)
	case openLastMsg:
		note := m.noteBook.LastNote()
		return m, openLinkCmd(note.ID)
	case noteBookLoadedMsg:
		m.noteBook = msg.noteBook
		// Force a rerender of the document
		if m.history.GetCurrent() != "" {
			note := m.noteBook.FindNote(m.history.GetCurrent())
			m.doc.renderNote(note)
		}
		return m, nil
	case openLinkMsg:
		m.openNote(msg.id, true)
		return m, nil
	case saveToCollectionMsg:
		m.noteBook.SaveToCollection(m.history.GetCurrent())
		return m, nil
	case resetCollectionMsg:
		m.noteBook.ResetCollection()
		return m, nil
	case tea.WindowSizeMsg:
		m.setSize(msg.Width, msg.Height)
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	// Render the entire UI

	var bar string
	if m.enteringCmd {
		bar = m.commandBar.View()
	} else {
		statusStyle := lipgloss.NewStyle().Width(m.width)
		bar = statusStyle.Render(m.statusMsg)
	}

	style := lipgloss.NewStyle().Foreground(colorDarkGray)

	idLen := core.IDLength + 5
	id := currentIdStyle.Render(m.doc.note.ID)
	vLine := style.Render(strings.Repeat("─", max(0, m.width-idLen))+
		"[ ") + id + style.Render(" ]"+"─")

	return lipgloss.JoinVertical(0, m.doc.View(), vLine, bar)
}

func (m *model) setSize(width, height int) {
	m.width, m.height = width, height

	m.doc.setSize(m.width, m.height-2)
	m.commandBar.setSize(m.width, 1)
}

func (m *model) openNote(id string, nav bool) {
	note := m.noteBook.FindNote(id)
	if note != nil {
		m.doc.renderNote(note)
		if nav {
			m.history.Push(id)
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
