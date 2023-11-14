package tui

import (
	"dnote/core"
	"dnote/render"
	"dnote/search"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type searchModel struct {
	width, height int

	keymap   docKeymap
	viewport viewport.Model

	query  string
	result *search.Result
	links  core.DocLinks
}

func newSearchModel() searchModel {
	return searchModel{0, 0, defaultDocKeyMap, viewport.New(0, 0), "", &search.Result{}, core.DocLinks{}}
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
		case m.links.GetLinkFromShortcut(msg.String()) != core.ShortcutLink{}:
			// Match any key that is a link shortcut
			return m, openLinkCmd(m.links.GetLinkFromShortcut(msg.String()).ID)
		}
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m searchModel) View() string {
	return m.viewport.View()
}

func (m *searchModel) render() {
	builder := new(strings.Builder)

	boxStyle := lipgloss.NewStyle().Margin(2, 1, 1).Width(m.width)
	style := lipgloss.NewStyle().Foreground(render.ColorHighCyan).Background(render.ColorLowBlue).PaddingLeft(1)
	queryStyle := lipgloss.NewStyle().Foreground(render.ColorWhite).Bold(true).Background(render.ColorLowBlue).PaddingRight(1)
	box := boxStyle.Render(fmt.Sprintf("%s%s",
		style.Render("Search results for: "),
		queryStyle.Render(m.query)))
	fmt.Fprint(builder, box)
	render.RenderLinkList(builder, m.result, &m.links, 0, render.DocLinkListStyles)
	m.viewport.SetContent(builder.String() + "\n")
}

func (m *searchModel) setSize(width, height int) {
	m.width, m.height = width, height
	m.viewport = viewport.New(width, height)
	m.render()
}

func (m *searchModel) setQuery(query string) {
	m.query = query
	m.render()
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
