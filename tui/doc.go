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
		key.WithKeys("o", "enter"),
		key.WithHelp("o or enter", "open link"),
	),
}

type openLinkMsg struct {
	id string
}

type cycleDirection uint

const (
	forward cycleDirection = iota
	backward
)

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

var shortcuts = []byte("ABCDEFGHIJKLMNOPQRSTUVXYZ")

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
			m.cycleLink(forward)
			return m, nil
		case key.Matches(msg, m.keymap.PrevLink):
			m.cycleLink(backward)
			return m, nil

		case key.Matches(msg, m.keymap.OpenLink):
			return m, openLinkCmd(m.src.links[m.selectedLink])
		case m.getShortCut(msg.String()) >= 0:
			return m, openLinkCmd(m.src.links[m.getShortCut(msg.String())])
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

func (m *docModel) cycleLink(dir cycleDirection) {
	length := len(m.src.links)

	if length <= 0 {
		return
	}

	if dir == forward {
		m.selectedLink = (m.selectedLink + 1) % length
	} else {
		m.selectedLink = (m.selectedLink + length - 1) % length
	}

	m.rerender()
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

	bracketStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#555555"))
	inactiveStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#aa00aa"))
	activeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffff55")).Bold(true)
	shortcutStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#55ff55"))
	re := regexp.MustCompile(fmt.Sprintf("\\|\\|([0-9]{%d})\\|\\|", core.IDLength))

	var idx = 0
	md = re.ReplaceAllStringFunc(md,
		func(s string) string {
			var style = inactiveStyle
			if idx == m.selectedLink {
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

func (m *docModel) getShortCut(s string) int {
	for i, ch := range shortcuts {
		if s == string(ch) && i < len(m.src.links) {
			return i
		}
	}

	return -1
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
