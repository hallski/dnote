package tui

import (
	"dnote/ext"
	"dnote/mdfiles"
	"dnote/render"
	"os"
	"time"

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
	noteBook  *mdfiles.MdDirectory
	gitStatus ext.GitStatus

	keymap appKeyMap

	history *history[historyItem]

	width  int
	height int

	showDoc bool
	doc     docModel
	search  searchModel

	enteringCmd bool
	commandBar  commandBar
}

func initialModel(noteBook *mdfiles.MdDirectory) model {
	return model{
		noteBook: noteBook,
		keymap:   defaultAppKeyMap,
		history:  NewHistory[historyItem](),

		showDoc:    true,
		doc:        newDoc(0, 0),
		search:     newSearchModel(noteBook),
		commandBar: newCommandBar(),
	}
}

func (m model) Init() tea.Cmd {
	return getGitStatusCmd(m.noteBook.Path())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		// If command bar is active, key bindings should be sent there
		case m.commandBar.inputActive():
			var cmd tea.Cmd
			m.commandBar, cmd = m.commandBar.Update(msg)
			return m, cmd
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Search):
			return m, emitMsgCmd(startSearchMsg{""})
		case key.Matches(msg, m.keymap.EditNode):
			return m, m.edit(false)
		case key.Matches(msg, m.keymap.EditNodeAlt):
			return m, m.edit(true)
		case key.Matches(msg, m.keymap.Back):
			m.setHistoryItem(m.history.GoBack())
			return m, nil
		case key.Matches(msg, m.keymap.Forward):
			m.setHistoryItem(m.history.GoForward())
			return m, nil
		case key.Matches(msg, m.keymap.RefreshNotes):
			return m, refreshNotebook(m.noteBook.Path())
		case key.Matches(msg, m.keymap.StartCmd):
			return m, m.commandBar.focus()
		case key.Matches(msg, m.keymap.QuickOpen):
			m.commandBar.startOpen(msg.String())
			return m, nil
		case key.Matches(msg, m.keymap.AddNote):
			return m, m.commandBar.startAdd(false)
		case key.Matches(msg, m.keymap.AddNoteAlt):
			return m, m.commandBar.startAdd(true)
		case key.Matches(msg, m.keymap.OpenRandomNote):
			return m, emitMsgCmd(openRandomMsg{})
		case key.Matches(msg, m.keymap.OpenLastNote):
			return m, emitMsgCmd(openLastMsg{})
		case key.Matches(msg, m.keymap.PrevNote):
			return m, emitMsgCmd(openPrevNoteMsg{})
		case key.Matches(msg, m.keymap.NextNote):
			return m, emitMsgCmd(openNextNoteMsg{})
		case key.Matches(msg, m.keymap.GitCommit):
			if m.gitStatus == ext.Dirty {
				return m, gitCommitCmd(m.noteBook.Path(), "")
			}
			return m, emitStatusMsgCmd("Nothing to commit")
		case key.Matches(msg, m.keymap.GitSync):
			return m, gitSyncCmd(m.noteBook.Path())
		}
		var cmd tea.Cmd
		if m.showDoc {
			m.doc, cmd = m.doc.Update(msg)
		} else {
			m.search, cmd = m.search.Update(msg)
		}
		return m, cmd
	case searchMsg:
		m.search.setQuery(msg.query)
		m.history.Push(historyItem{kindSearch, msg.query})
		m.showDoc = false
	case startSearchMsg:
		m.commandBar.startSearch(msg.query)
	case gitCommandFinishedMsg:
		return m, tea.Batch(
			emitStatusMsgCmd(msg.result), getGitStatusCmd(m.noteBook.Path()),
		)
	case notesDirModifiedMsg:
		path := m.noteBook.Path()
		return m, tea.Batch(refreshNotebook(path), getGitStatusCmd(path))
	case refreshNotebookMsg:
		return m, refreshNotebook(m.noteBook.Path())
	case gitStatusMsg:
		m.gitStatus = msg.status
		return m, nil
	case openRandomMsg:
		note := m.noteBook.RandomNote()
		return m, openLinkCmd(note.ID)
	case openPrevNoteMsg:
		if m.showDoc {
			return m, openInDirectionCmd(m.noteBook, m.doc.note, mdfiles.Backward)
		}
	case openNextNoteMsg:
		if m.showDoc {
			return m, openInDirectionCmd(m.noteBook, m.doc.note, mdfiles.Forward)
		}
	case addNoteMessage:
		note, err := m.noteBook.CreateNote(msg.title)
		if err != nil {
			return m, emitStatusMsgCmd("Error creating note")
		}
		return m, openEditor(m.noteBook, note.ID, msg.keepFocus)
	case openLastMsg:
		note := m.noteBook.LastNote()
		return m, openLinkCmd(note.ID)
	case noteBookLoadedMsg:
		m.setNotebook(msg.noteBook)
		return m, emitStatusMsgCmd("Refreshed")
	case openLinkMsg:
		m.openNote(msg.id, true)
	case saveToCollectionMsg:
		item := m.history.GetCurrent()

		if item != (historyItem{}) && item.kind == kindNote {
			m.noteBook.SaveToCollection(item.value)
		}
	case resetCollectionMsg:
		m.noteBook.ResetCollection()
	case tea.WindowSizeMsg:
		m.setSize(msg.Width, msg.Height)
	}

	var cbCmd, viewCmd tea.Cmd
	m.commandBar, cbCmd = m.commandBar.Update(msg)
	if m.showDoc {
		m.doc, viewCmd = m.doc.Update(msg)
	} else {
		m.search, viewCmd = m.search.Update(msg)
	}
	return m, tea.Batch(cbCmd, viewCmd)
}

// Render the entire UI
func (m model) View() string {
	title := render.Titlebar(m.width, m.noteBook.LastNote().ID, m.gitStatus)

	view := ""
	if m.showDoc {
		view = m.doc.View()
	} else {
		view = m.search.View()
	}

	return lipgloss.JoinVertical(0, title, view, m.commandBar.View())
}

func (m *model) setSize(width, height int) {
	const totalBarSize = render.TitleBarHeight + CommandBarHeight

	m.width, m.height = width, height

	mainViewSize := m.height - totalBarSize

	m.doc.setSize(m.width, mainViewSize)
	m.search.setSize(m.width, mainViewSize)
	m.commandBar.setSize(m.width, CommandBarHeight)
}

func (m *model) setNotebook(notebook *mdfiles.MdDirectory) {
	m.noteBook = notebook
	m.search.setCollection(notebook)

	// Force a rerender of the current document
	if m.history.curPos >= 0 {
		m.setHistoryItem(m.history.GetCurrent())
	}
}

func (m *model) edit(keepFocus bool) tea.Cmd {
	item := m.history.GetCurrent()
	if item.kind != kindNote {
		return nil
	}

	return openEditor(m.noteBook, item.value, keepFocus)
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
	}
}

// Poor mans directory watcher
// Tried with fsnotify/fsnotify but it gave lots of events for each
// change and stopped triggering after a short while
func startFileMonitor(p *tea.Program, path string) {
	go func() {
		fileInfo, err := os.Stat(path)
		if err != nil {
			panic(err)
		}

		for {
			time.Sleep(500 * time.Millisecond)

			fi, err := os.Stat(path)
			if err != nil {
				panic(err)
			}

			if fi.ModTime() != fileInfo.ModTime() {
				p.Send(notesDirModifiedMsg{})
			}
			fileInfo = fi
		}
	}()
}

func Run(noteBook *mdfiles.MdDirectory, openId string) error {
	m := initialModel(noteBook)
	m.openNote(openId, true)
	p := tea.NewProgram(m)

	startFileMonitor(p, noteBook.Path())

	_, err := p.Run()

	return err
}
