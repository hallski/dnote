package tui

import (
	"dnote/core"
	"dnote/mdfiles"
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type selectedLink struct {
	ID    string
	index int
}

type docModel struct {
	keymap docKeymap

	links docLinks

	width  int
	height int

	note         *core.Note
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
			// Match any key that is a link shortcut
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

func (m *docModel) nextLink() {
	m.links.Next()
	m.render()
}

func (m *docModel) prevLink() {
	m.links.Prev()
	m.render()
}

var linkReplacementRE = regexp.MustCompile(fmt.Sprintf("\\|\\|([0-9]{%d})\\|\\|",
	core.IDLength))

func (m *docModel) processNoteContent() {
	var links []string
	processed := mdfiles.LinkRegexp.ReplaceAllStringFunc(m.note.Content,
		func(s string) string {
			id := s[2:5]
			links = append(links, id)
			return "||" + id + "||"
		},
	)

	for _, bl := range m.note.BackLinks {
		links = append(links, bl.ID)
	}

	m.links = newDocLinks(links)
	m.preprocessed = processed
}

func (m *docModel) render() {
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(m.width),
	)

	md, err := r.Render(m.preprocessed)
	if err != nil {
		panic(err)
	}

	var idx = 0
	md = linkReplacementRE.ReplaceAllStringFunc(md,
		func(s string) string {
			active := m.links.IsActive(idx)
			sc := m.links.GetShortcut(idx)
			idx++
			return renderLink(s[2:5], sc, active, docLinkStyles)
		},
	)

	// Crude backlink support
	builder := new(strings.Builder)
	if len(m.note.BackLinks) > 0 {
		bls := new(strings.Builder)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("#5555ff")).Background(lipgloss.Color("#222222"))

		fmt.Fprintln(bls, backlinksTitleStyle.Render("Backlinks"))

		for i, bl := range m.note.BackLinks {
			linkIdx := i + idx
			link := m.links.GetLinkIdx(linkIdx)
			active := m.links.IsActive(linkIdx)
			sc := m.links.GetShortcut(linkIdx)
			fmt.Fprintf(bls, "%s%s\n",
				renderLink(link, sc, active, backLinkStyles),
				style.Render(" "+bl.Title))
		}

		box := backlinksBoxStyle.Copy().
			Width(m.width - backlinksBoxStyle.GetHorizontalBorderSize())

		fmt.Fprintf(builder, box.Render(bls.String()))
	}

	fmt.Fprint(builder, docNoteIdStyle.Width(m.width).Render("#"+m.note.ID))

	m.viewport.SetContent(md + "\n" + builder.String())
}

func renderLink(link, sc string, active bool, styles linkStyles) string {
	var style = styles.inactive
	if active {
		style = styles.active
	}

	if sc == "" {
		return style.Render(styles.bracket.Render("[[") +
			style.Render(link) +
			styles.bracket.Render("]]"),
		)
	}

	return style.Render(styles.bracket.Render("[") +
		styles.shortcut.Render(sc) +
		styles.bracket.Render("|") +
		style.Render(link) +
		styles.bracket.Render("]"),
	)
}

func (m *docModel) renderNote(note *core.Note) {
	m.note = note
	m.processNoteContent()
	m.render()
	m.viewport.SetYOffset(0)
}

func (m *docModel) setSize(width, height int) {
	m.width = width
	m.height = height
	m.viewport = viewport.New(width, height)
	m.render()
}

func newDoc(width, height int) docModel {
	m := docModel{
		keymap:   DefaultDocKeyMap,
		viewport: viewport.New(width, height),
		links:    newDocLinks([]string{}),
	}

	return m
}
