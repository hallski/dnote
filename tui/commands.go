package tui

import (
	"dnote/core"
	"dnote/ext"
	"dnote/mdfiles"
	"fmt"
	"time"

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

func startEditor(path string, keepFocus bool) tea.Cmd {
	return func() tea.Msg {
		c := ext.GetEditorNewPane(path, keepFocus)
		if err := c.Run(); err != nil {
			return statusMsg{fmt.Sprintf("Failed editing: %s", err)}
		}

		return editorStartedMsg{}
	}
}

func openEditor(noteBook *mdfiles.MdDirectory, id string, keepFocus bool) tea.Cmd {
	note := noteBook.FindNote(id)
	if note == nil {
		return func() tea.Msg { return statusMsg{"Failed opening " + id} }
	}

	return startEditor(note.Path, keepFocus)
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

func addNoteCmd(title string, keepFocus bool) tea.Cmd {
	return func() tea.Msg {
		return addNoteMessage{title, keepFocus}
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

func timeoutStatusCmd(statusId int) tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return statusMsgTimeoutMsg{statusId}
	})
}

func getGitStatusCmd(path string) tea.Cmd {
	return func() tea.Msg {
		client, err := ext.NewGitClient(path)
		if err != nil {
			return statusMsg{"Couldn't find Git executable"}
		}

		status, err := client.Status()
		if err != nil {
			return statusMsg{fmt.Sprintf("Failed to check git status: %s", err)}
		}

		return gitStatusMsg{status}
	}
}

func gitCommitCmd(path string, msg string) tea.Cmd {
	if msg == "" {
		msg = "Update from dNote"
	}

	return func() tea.Msg {
		client, err := ext.NewGitClient(path)
		if err != nil {
			return statusMsg{"Couldn't find Git executable"}
		}

		if err := client.Commit(msg); err != nil {
			return statusMsg{fmt.Sprintf("Failed to commit: %s", err)}
		}

		return gitCommandFinishedMsg{"Commited"}
	}
}

func gitSyncCmd(path string) tea.Cmd {
	return func() tea.Msg {
		client, err := ext.NewGitClient(path)
		if err != nil {
			return statusMsg{"Couldn't find Git executable"}
		}

		if err := client.PullRebasePush(); err != nil {
			return statusMsg{"Failed to sync with remote"}
		}

		return gitCommandFinishedMsg{"Successfully sycned with remote"}
	}
}
