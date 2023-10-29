package tui

import (
	"hallski/thinkadus/core"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	screen      tea.Model
	path        string
	notes       []core.Note
	currentNote *core.Note
	width       int
	height      int
}

func initialUIModel(path string) Model {
	return Model{
		screen:      InitLoadingScreen(),
		path:        path,
		notes:       nil,
		currentNote: nil,
		width:       0,
		height:      0,
	}
}
func (m Model) Init() tea.Cmd {
	return tea.Batch(ReadFilesCmd(m.path), m.screen.Init())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case notesMsg:
		m.notes = msg.notes
		return m, NavigateToSearch()

	case navigateToSearchMsg:
		m.screen = InitListNotesScreen(m.notes, m.width, m.height)
		return m, nil

	case navigateToNoteMsg:
		//note := msg.note
		//fmt.Println("Navigate to note", note.Path)
		m.screen = InitNoteScreen(msg.note, m.width, m.height)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			m.screen, cmd = m.screen.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.screen, cmd = m.screen.Update(msg)
		return m, nil
	default:
		m.screen, cmd = m.screen.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	return m.screen.View()
}

func StartTui(path string) {
	p := initialUIModel(path)
	prompt := tea.NewProgram(p)
	if err := prompt.Start(); err != nil {
		panic(err)
	}
}
