package tui

import (
	"hallski/thinkadus/core"
	"hallski/thinkadus/shell"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// Messages
type navigateToSearchMsg struct{}
type navigateToNoteMsg struct{ note core.Note }
type notesMsg struct {
	notes []core.Note
}

func NavigateToSearch() tea.Cmd {
	return func() tea.Msg {
		return navigateToSearchMsg{}
	}
}

func NavigateToNote(note core.Note) tea.Cmd {
	return func() tea.Msg {
		return navigateToNoteMsg{note}
	}
}

func ReadFilesCmd(path string) tea.Cmd {
	return func() tea.Msg {
		notes, err := shell.ReadFiles(path)

		if err != nil {
			log.Fatal(err)
		}

		return notesMsg{notes}
	}
}
