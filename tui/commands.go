package tui

import (
	"dnote/core"
	"dnote/ext"
	"dnote/mdfiles"
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// Command to send a message
// Used by sub views to pass messages back to the program
func emitMsgCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func emitStatusMsgCmd(msg string) tea.Cmd {
	return func() tea.Msg {
		return statusMsg{msg}
	}
}

func openLinkCmd(id string) tea.Cmd {
	return func() tea.Msg {
		return openLinkMsg{id}
	}
}

func startEditor(path string, newPane bool) tea.Cmd {
	if newPane {

		return func() tea.Msg {
			c := ext.GetEditorNewPane(path)

			if err := c.Run(); err != nil {
				return statusMsg{fmt.Sprintf("Failed editing: %s", err)}
			}

			return editorFinishedMsg{}
		}
	}

	editor := ext.GetEditor()

	c := exec.Command(editor, path)

	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			return statusMsg{fmt.Sprintf("Failed editing: %s", err)}
		} else {
			return editorFinishedMsg{}
		}
	})
}

func openEditor(noteBook *mdfiles.MdDirectory, id string, newPane bool) tea.Cmd {
	note := noteBook.FindNote(id)
	if note == nil {
		return func() tea.Msg { return statusMsg{"Failed opening " + id} }
	}

	return startEditor(note.Path, newPane)
}

func refreshNotebook(path string) tea.Cmd {
	return func() tea.Msg {
		noteBook, err := mdfiles.Load(path)
		if err != nil {
			panic(err)
		}

		return noteBookLoadedMsg{noteBook}
	}
}

func addNoteCmd(title string, newPane bool) tea.Cmd {
	return func() tea.Msg {
		return addNoteMessage{title, newPane}
	}
}

func searchCmd(query string) tea.Cmd {
	return func() tea.Msg {
		return searchMsg{query}
	}
}

func openInDirectionCmd(notebook *mdfiles.MdDirectory, note *core.Note, dir mdfiles.CycleDirection) tea.Cmd {
	newNote := notebook.NoteInDirection(note, dir)
	if newNote == nil {
		return nil
	}

	return openLinkCmd(newNote.ID)
}
