package tui

import (
	"dnote/mdfiles"
	"dnote/render"
	"dnote/search"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type historyKind uint

const (
	kindNote historyKind = iota
	kindSearch
)

type historyItem struct {
	kind  historyKind
	value string
}

type model struct {
	noteBook *mdfiles.MdDirectory

	keymap appKeyMap

	statusMsg string

	history *history[historyItem]

	width  int
	height int

	showDoc bool
	doc     docModel

	searchResult *search.Result
	search       searchModel

	enteringCmd bool
	commandBar  commandBar
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	return model{
		noteBook,
		defaultAppKeyMap,
		"",
		NewHistory[historyItem](),
		0, 0,
		true,
		newDoc(0, 0),
		nil,
		newSearchModel(),
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
		case m.enteringCmd:
			var cmd tea.Cmd
			m.commandBar, cmd = m.commandBar.Update(msg)
			return m, cmd
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Search):
			m.commandBar.startSearch("")
			m.enteringCmd = true
			return m, nil
		case key.Matches(msg, m.keymap.EditNode):
			return m, m.edit()
		case key.Matches(msg, m.keymap.Back):
			m.setHistoryItem(m.history.GoBack())
			return m, nil
		case key.Matches(msg, m.keymap.Forward):
			m.setHistoryItem(m.history.GoForward())
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
		if m.showDoc {
			m.doc, cmd = m.doc.Update(msg)
		} else {
			m.search, cmd = m.search.Update(msg)
		}
		return m, cmd
	case searchMsg:
		m.search.setQuery(msg.search)
		m.searchResult = search.NewFullText(msg.search, m.noteBook)
		m.search.setResult(m.searchResult)
		m.history.Push(historyItem{kindSearch, msg.search})
		m.showDoc = false
		return m, nil
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
		if m.history.curPos >= 0 {
			m.setHistoryItem(m.history.GetCurrent())
		}
		return m, nil
	case openLinkMsg:
		m.openNote(msg.id, true)
		return m, nil
	case saveToCollectionMsg:
		item := m.history.GetCurrent()

		if item != (historyItem{}) && item.kind == kindNote {
			m.noteBook.SaveToCollection(item.value)
		}
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

	bottomBar := ""
	if m.showDoc {
		bottomBar = render.BottomBarNote(m.doc.note, m.width)
	} else {
		bottomBar = render.BottomBarSearch(m.searchResult, m.width)
	}

	title := render.Titlebar(m.width, m.noteBook.LastNote().ID)
	if m.showDoc {
		return lipgloss.JoinVertical(0, title, m.doc.View(), bottomBar, bar)
	}

	return lipgloss.JoinVertical(0, title, m.search.View(), bottomBar, bar)
}

func (m *model) setSize(width, height int) {
	m.width, m.height = width, height

	m.doc.setSize(m.width, m.height-5)
	m.search.setSize(m.width, m.height-5)
	m.commandBar.setSize(m.width, 1)
}

func (m *model) edit() tea.Cmd {
	item := m.history.GetCurrent()
	if item.kind != kindNote {
		return nil
	}

	return openEditor(m.noteBook, item.value)
}

func (m *model) openNote(id string, nav bool) {
	note := m.noteBook.FindNote(id)
	if note != nil {
		m.doc.renderNote(note)
		if nav {
			m.history.Push(historyItem{kindNote, note.ID})
		}
	}

	m.showDoc = true
}

func (m *model) setHistoryItem(item historyItem) {
	switch item.kind {
	case kindNote:
		m.openNote(item.value, false)
	case kindSearch:
		m.showDoc = false
		m.search.setQuery(item.value)
		m.search.setResult(search.NewFullText(item.value, m.noteBook))
	}
}

func Run(noteBook *mdfiles.MdDirectory, openId string) error {
	m := initialModel(noteBook)
	m.openNote(openId, true)
	p := tea.NewProgram(m)

	_, err := p.Run()

	return err
}
