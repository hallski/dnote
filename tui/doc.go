package tui

import (
	"dnote/core"
	"dnote/mdfiles"
	"fmt"
	"regexp"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type docKeymap struct {
	NextLink key.Binding
	PrevLink key.Binding
	OpenLink key.Binding
}

var DefaultDocKeyMap = docKeymap{
	NextLink: key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl+n", "next link"),
	),
	PrevLink: key.NewBinding(
		key.WithKeys("ctrl+p"),
		key.WithHelp("ctrl+p", "prev link"),
	),
	OpenLink: key.NewBinding(
		key.WithKeys("ctrl+o"),
		key.WithHelp("ctrl+o", "open link"),
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
	ID    string
	index int
}

type docModel struct {
	keymap docKeymap

	links docLinks

	width  int
	height int

	preprocessed string

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
		case m.links.GetLink(msg.String()) != "":
			return m, openLinkCmd(m.links.GetLink(msg.String()))
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

func (m *docModel) processNoteContent(content string) {
	var links []string
	processed := mdfiles.LinkRegexp.ReplaceAllStringFunc(content,
		func(s string) string {
			id := s[2:5]
			links = append(links, id)
			return "||" + id + "||"
		},
	)

	m.links = newDocLinks(links)
	m.preprocessed = processed
}

func (m *docModel) nextLink() {
	m.links.Next()
	m.rerender()
}

func (m *docModel) prevLink() {
	m.links.Prev()
	m.rerender()
}

func (m *docModel) rerender() {
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(m.width),
	)

	md, err := r.Render(m.preprocessed)
	if err != nil {
		panic(err)
	}

	bracketStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	inactiveStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#aa00aa"))
	activeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff55")).Bold(true)
	shortcutStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#55ff55"))
	re := regexp.MustCompile(fmt.Sprintf("\\|\\|([0-9]{%d})\\|\\|", core.IDLength))

	var idx = 0
	md = re.ReplaceAllStringFunc(md,
		func(s string) string {
			var style = inactiveStyle
			if m.links.IsActive(idx) {
				style = activeStyle
			}

			sc := string(shortcuts[idx])
			idx++

			return style.Render(bracketStyle.Render("[") +
				shortcutStyle.Render(sc) +
				bracketStyle.Render("|") +
				style.Render(s[2:5]) +
				bracketStyle.Render("]"),
			)
		},
	)

	m.viewport.SetContent(md)
}

func (m *docModel) renderNote(note *core.Note) {
	m.processNoteContent(note.Content)
	m.rerender()
}

func (m *docModel) setSize(width, height int) {
	m.viewport = viewport.New(width, height)
	m.rerender()
}

func newDoc(width, height int) docModel {
	m := docModel{
		keymap:   DefaultDocKeyMap,
		viewport: viewport.New(width, height),
		links:    newDocLinks([]string{}),
	}

	return m
}
