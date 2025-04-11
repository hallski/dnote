package tui

import (
	"dnote/core"
	"dnote/tui/render"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type docModel struct {
	keymap docKeymap

	links core.DocLinks

	width, height int

	note *core.Note

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
			m.nextLink()
			return m, nil
		case key.Matches(msg, m.keymap.PrevLink):
			m.prevLink()
			return m, nil

		case key.Matches(msg, m.keymap.OpenLink):
			l := m.links.Current()
			if l != "" {
				return m, openLinkCmd(l)
			}
		case m.links.HasShortcut(msg.String()):
			// Match any key that is a link shortcut
			l := m.links.GetLinkFromShortcut(msg.String())
			return m, openLinkCmd(l.ID)
		case m.altShortcut(msg.String()) != core.ShortcutLink{}:
			link := m.altShortcut(msg.String())
			return m, emitMsgCmd(openEditorWithNoteIdMsg{link.ID, true})
		}
	}
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m docModel) View() string {
	bottomBar := render.BottomBarNote(m.note, m.width)

	s := lipgloss.NewStyle().PaddingLeft(1)
	return lipgloss.JoinVertical(0, s.Render(m.viewport.View()), bottomBar)
}

func (m *docModel) nextLink() {
	m.links.Next()
	m.render()
}

func (m *docModel) prevLink() {
	m.links.Prev()
	m.render()
}

func (m *docModel) render() {
	md, idx := render.Note(m.note, &m.links, m.width)
	backlinks := render.BackLinks(m.note, idx, &m.links, m.width)

	m.viewport.SetContent(md + "\n" + backlinks + "\n")
}

func (m *docModel) renderNote(note *core.Note) {
	m.note = note
	m.render()
	m.viewport.SetYOffset(0)
}

func (m *docModel) setSize(width, height int) {
	oldWidth := m.width
	m.width, m.height = width, height

	if oldWidth != width {
		m.render()
	}

	m.viewport.Width = width
	m.viewport.Height = height - render.BottomBarHeight
}

func (m *docModel) altShortcut(keys string) core.ShortcutLink {
	if !strings.HasPrefix(keys, "alt+") {
		return core.ShortcutLink{}
	}

	return m.links.GetLinkFromShortcut(keys[4:])
}

func newDoc(width, height int) docModel {
	m := docModel{
		keymap:   defaultDocKeyMap,
		viewport: viewport.New(width, height),
		links:    core.NewDocLinks([]string{}),
	}

	return m
}
