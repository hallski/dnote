package tui

import (
	"dnote/core"
	"dnote/search"
	"dnote/tui/render"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type searchModel struct {
	width, height int

	collection search.FullTextCollection

	keymap   searchKeymap
	viewport viewport.Model

	result *search.Result
	links  core.DocLinks

	showTags bool
}

func newSearchModel(collection search.FullTextCollection) searchModel {
	return searchModel{collection: collection, keymap: defaultSearchKeyMap}
}

func (m searchModel) Init() tea.Cmd {
	return nil
}

func (m searchModel) Update(msg tea.Msg) (searchModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.NextLink):
			m.links.Next()
			m.render()
			return m, nil
		case key.Matches(msg, m.keymap.PrevLink):
			m.links.Prev()
			m.render()
			return m, nil

		case key.Matches(msg, m.keymap.OpenLink):
			l := m.links.Current()
			if l != "" {
				return m, openLinkCmd(l)
			}
		case key.Matches(msg, m.keymap.ExtendSearch):
			return m, emitMsgCmd(startSearchMsg{m.result.Query})
		case key.Matches(msg, m.keymap.ToggleTags):
			m.showTags = !m.showTags
			m.render()
			return m, nil
		case m.links.GetLinkFromShortcut(msg.String()) != core.ShortcutLink{}:
			// Match any key that is a link shortcut
			return m, openLinkCmd(m.links.GetLinkFromShortcut(msg.String()).ID)
		case m.altShortcut(msg.String()) != core.ShortcutLink{}:
			link := m.altShortcut(msg.String())
			return m, emitMsgCmd(openEditorWithNoteIdMsg{link.ID, true})
		}

		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m searchModel) View() string {
	bottomBar := render.BottomBarSearch(m.result, m.width)
	s := lipgloss.NewStyle().PaddingLeft(1)

	return lipgloss.JoinVertical(0, s.Render(m.viewport.View()), bottomBar)
}

var style = lipgloss.NewStyle().
	Foreground(render.ColorWhite).
	Padding(0, 1).
	Underline(true).
	MarginRight(1).
	Bold(true)

var queryStyle = lipgloss.NewStyle().
	Foreground(render.ColorYellow).
	Bold(true).
	PaddingRight(1)

func (m *searchModel) render() {
	if m.result == nil {
		return
	}
	builder := new(strings.Builder)

	boxStyle := lipgloss.NewStyle().Margin(1, 2, 1).Width(m.width)
	var box = boxStyle.Render(fmt.Sprintf("%s%s",
		style.Render("Search results for:"),
		queryStyle.Render(m.result.Query)))

	fmt.Fprintln(builder, box)
	render.LinkList(builder,
		m.result, &m.links, 0, m.showTags, render.DocLinkListStyles)

	m.viewport.SetContent(builder.String() + "\n")
}

func (m *searchModel) setSize(width, height int) {
	m.width, m.height = width, height
	m.viewport = viewport.New(width, height-render.BottomBarHeight)
	m.render()
}

func (m *searchModel) setCollection(collection search.FullTextCollection) {
	m.collection = collection
}

func (m *searchModel) setQuery(query string) {
	m.setResult(search.NewFullText(query, m.collection))
}

func (m *searchModel) setResult(result *search.Result) {
	m.result = result
	var ll []string
	for _, note := range m.result.ListNotes() {
		ll = append(ll, note.ID)
	}

	m.links = core.NewDocLinks(ll)
	m.render()
}

func (m *searchModel) altShortcut(keys string) core.ShortcutLink {
	if !strings.HasPrefix(keys, "alt+") {
		return core.ShortcutLink{}
	}

	return m.links.GetLinkFromShortcut(keys[4:])
}
