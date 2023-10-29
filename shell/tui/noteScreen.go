package tui

import (
	"fmt"
	"hallski/thinkadus/core"
	"hallski/thinkadus/shell"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type NoteModel struct {
	note     core.Note
	viewPort viewport.Model
}

func (m NoteModel) Init() tea.Cmd {
	// Read content of note
	return nil
}

func (m NoteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "o":
			return m, NavigateToSearch()
		default:
			var cmd tea.Cmd
			m.viewPort, cmd = m.viewPort.Update(msg)
			return m, cmd
		}
	default:
		var cmd tea.Cmd
		m.viewPort, cmd = m.viewPort.Update(msg)
		return m, cmd
	}
}

func (m NoteModel) View() string {
	return m.viewPort.View()
}

func InitNoteScreen(note core.Note, width, height int) tea.Model {
	fmt.Println("Init note screen", note.Path)
	viewPort := viewport.New(width, height)

	content, err := shell.ReadNoteContent(note.Path)
	if err != nil {
		panic(err)
	}

	md, err := glamour.Render(content, "dark")
	if err != nil {
		panic(err)
	}

	viewPort.SetContent(md)

	return NoteModel{note, viewPort}
}
