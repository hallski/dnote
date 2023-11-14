package tui

import (
	"dnote/core"
	"dnote/render"
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

type selectedLink struct {
	ID    string
	index int
}

type docModel struct {
	keymap docKeymap

	links core.DocLinks

	width, height int

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

// Adds a hack to support wiki links even though Glamour do not support them
// Since [[ is part of ANSI escape codes, replace them with || before parsing
// with Glamour.
// The *qwq* before and after is to ensure that Glamor will insert separate
// style codes for those (effectively, ensuring that whatever comes after the
// link will get a new style applied.
func addLinkHack(id string) string {
	return "||" + id + "||*qwq*"
}

// Remove the insert qwq (only leaving the escape code and reset in the document
// This is fine as nothing will actually use those codes
func removeLinkStyleHack(s string) string {
	return strings.Replace(s, "qwq", "", -1)
}

func (m *docModel) processNoteContent() {
	var links []string
	processed := core.LinkRegexp.ReplaceAllStringFunc(m.note.Content,
		func(s string) string {
			id := s[2:5]
			links = append(links, id)
			return addLinkHack(id)
		},
	)

	for _, bl := range m.note.BackLinks.ListNotes() {
		links = append(links, bl.ID)
	}

	m.links = core.NewDocLinks(links)
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

	md = removeLinkStyleHack(md)

	var idx = 0
	md = linkReplacementRE.ReplaceAllStringFunc(md,
		func(s string) string {
			linkID := s[2:5]
			active := m.links.IsActive(idx)
			sc := m.links.GetShortcut(linkID)
			idx++

			return renderLink(linkID, sc, active, render.DocLinkStyles)
		},
	)

	// Crude backlink support
	builder := new(strings.Builder)
	if len(m.note.BackLinks.ListNotes()) > 0 {
		bls := new(strings.Builder)

		beforeText := "─ Backlinks "
		beforeLen := m.width - len([]rune(beforeText))
		if beforeLen > 0 {
			border := strings.Repeat("─", beforeLen)
			fmt.Fprintln(bls, render.BacklinksTitleStyle.Render(beforeText+border+"\n"))
		}

		render.RenderLinkList(bls, m.note.BackLinks, &m.links, idx)

		box := render.BacklinksBoxStyle.Copy().
			Width(m.width - render.BacklinksBoxStyle.GetHorizontalBorderSize())

		fmt.Fprintf(builder, box.Render(bls.String()))
	}

	m.viewport.SetContent(md + "\n" + builder.String() + "\n")
}

func renderLink(link, sc string, active bool, styles render.LinkStyles) string {
	var style = styles.Inactive
	if active {
		style = styles.Active
	}

	if sc == "" {
		return styles.Bracket.Render("[[") +
			style.Render(link) +
			styles.Bracket.Render("]]")
	}

	return styles.Bracket.Render("[") +
		styles.Shortcut.Render(sc) +
		styles.Bracket.Render("|") +
		style.Render(link) +
		styles.Bracket.Render("]")
}

func (m *docModel) renderNote(note *core.Note) {
	m.note = note
	m.processNoteContent()
	m.render()
	m.viewport.SetYOffset(0)
}

func (m *docModel) setSize(width, height int) {
	m.width, m.height = width, height
	m.viewport = viewport.New(width, height)
	m.render()
}

func newDoc(width, height int) docModel {
	m := docModel{
		keymap:   defaultDocKeyMap,
		viewport: viewport.New(width, height),
		links:    core.NewDocLinks([]string{}),
	}

	return m
}
