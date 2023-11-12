package tui

import "dnote/mdfiles"

type exitCmdMsg struct{}
type openLinkMsg struct {
	id string
}

type statusMsg struct{ s string }
type editorFinishedMsg struct{}
type refreshNotebookMsg struct{}
type noteBookLoadedMsg struct{ noteBook *mdfiles.MdDirectory }

type openRandomMsg struct{}
type openLastMsg struct{}

type resetCollectionMsg struct{}
type saveToCollectionMsg struct{}
