package tui

import (
	"dnote/core"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type docKeymap struct {
	NextLink key.Binding
	OpenLink key.Binding
}

var DefaultDocKeyMap = docKeymap{
	NextLink: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next link"),
	),
	OpenLink: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "open link"),
	),
}

type openLinkMsg struct {
	id string
}

func openLinkCmd(id string) tea.Cmd {
	return func() tea.Msg {
		return openLinkMsg{id}
	}
}

type selectedLink struct {
	ID string
}

type docModel struct {
	keymap docKeymap

	width  int
	height int

	renderedMd string

	selectedLink selectedLink

	viewport viewport.Model
}

func (m docModel) Init() tea.Cmd {
	return nil
}

func (m docModel) Update(msg tea.Msg) (docModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.NextLink):
			// Move to next link
		case key.Matches(msg, m.keymap.OpenLink):
			return m, openLinkCmd(m.selectedLink.ID)
		}
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m docModel) View() string {
	return m.viewport.View()
}

func (m *docModel) renderNote(note *core.Note) {
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(m.width),
	)
	md, err := r.Render(note.Content)
	if err != nil {
		panic(err)
	}

	m.renderedMd = md
	m.viewport.SetContent(md)
}

func (m *docModel) setSize(width, height int) {
	m.viewport = viewport.New(width, height)
	m.viewport.SetContent(m.renderedMd)
}

func newDoc(width, height int, note *core.Note) docModel {
	m := docModel{
		keymap:       DefaultDocKeyMap,
		viewport:     viewport.New(width, height),
		selectedLink: selectedLink{ID: "1337"},
	}

	if note != nil {
		fmt.Printf("Rendering note: %s\n", note.Title)
		m.renderNote(note)
	}

	return m
}
