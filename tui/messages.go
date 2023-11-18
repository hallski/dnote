package tui

import (
	"dnote/ext"
	"dnote/mdfiles"
)

type openLinkMsg struct {
	id string
}

type statusMsg struct{ s string }
type statusMsgTimeoutMsg struct{ id int }
type editorStartedMsg struct{}
type notesDirModifiedMsg struct{}
type refreshNotebookMsg struct{}
type noteBookLoadedMsg struct{ noteBook *mdfiles.MdDirectory }

type openRandomMsg struct{}
type openLastMsg struct{}
type openNextNoteMsg struct{}
type openPrevNoteMsg struct{}

type addNoteMessage struct {
	title     string
	keepFocus bool
}

type resetCollectionMsg struct{}
type saveToCollectionMsg struct{}

type startSearchMsg struct{ query string }
type searchMsg struct{ query string }

type gitStatusMsg struct{ status ext.GitStatus }
type gitCommandFinishedMsg struct{ result string }
