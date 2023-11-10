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

type preparedSource struct {
	links        []string
	preprocessed string
}
type selectedLink struct {
	ID    string
	index int
}

type docModel struct {
	keymap docKeymap

	width  int
	height int

	src          preparedSource
	selectedLink int

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
			if len(m.src.links) > 0 {
				m.selectedLink = (m.selectedLink + 1) % len(m.src.links)
			} else {
				m.selectedLink = -1
			}
			m.rerender()
			return m, nil
		case key.Matches(msg, m.keymap.OpenLink):
			return m, openLinkCmd(m.src.links[m.selectedLink])
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

func processNoteContent(content string) preparedSource {
	var links []string
	processed := mdfiles.LinkRegexp.ReplaceAllStringFunc(content,
		func(s string) string {
			id := s[2:5]
			links = append(links, id)
			return "||" + id + "||"
		},
	)

	return preparedSource{links, processed}

}

func (m *docModel) rerender() {
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(m.width),
	)

	md, err := r.Render(m.src.preprocessed)
	if err != nil {
		panic(err)
	}

	inactiveStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff00ff"))
	activeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00ffff"))
	re := regexp.MustCompile(fmt.Sprintf("\\|\\|([0-9]{%d})\\|\\|", core.IDLength))

	var idx = 0
	md = re.ReplaceAllStringFunc(md,
		func(s string) string {
			var style = inactiveStyle
			if idx == m.selectedLink {
				style = activeStyle
			}

			idx++

			return style.Render("[[" + s[2:5] + "]]")
		},
	)

	m.viewport.SetContent(md)
}

func (m *docModel) renderNote(note *core.Note) {
	m.src = processNoteContent(note.Content)
	m.selectedLink = -1

	m.rerender()
}

func (m *docModel) setSize(width, height int) {
	m.viewport = viewport.New(width, height)
	m.rerender()
}

func newDoc(width, height int) docModel {
	m := docModel{
		keymap:       DefaultDocKeyMap,
		viewport:     viewport.New(width, height),
		selectedLink: -1,
	}

	return m
}
