package render

import (
	"dnote/core"
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/glamour"
)

const docMaxWidth = 84

var linkReplacementRE = regexp.MustCompile(fmt.Sprintf("\\|\\|([0-9]{%d})\\|\\|",
	core.IDLength))

var tagsRE = regexp.MustCompile("\\s(#+[-\\w]+)")

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

func Note(note *core.Note, docLinks *core.DocLinks, width int) (string, int) {
	r, err := glamour.NewTermRenderer(
		glamour.WithStyles(GetGlamming()),
		glamour.WithWordWrap(min(width, docMaxWidth)),
	)

	var links []string
	processed := core.LinkRegexp.ReplaceAllStringFunc(note.Content,
		func(s string) string {
			id := s[2:5]
			links = append(links, id)
			return addLinkHack(id)
		},
	)

	re := regexp.MustCompile("[^# \\n]")
	reTagBack := regexp.MustCompile("#+\\$+")

	var tags []string
	processed = tagsRE.ReplaceAllStringFunc(processed,
		func(s string) string {
			tag := s[1:]
			tags = append(tags, tag)

			replacement := re.ReplaceAllString(s, "$")
			return replacement
		},
	)

	for _, bl := range note.BackLinks.ListNotes() {
		links = append(links, bl.ID)
	}

	*docLinks = core.NewDocLinks(links)
	preprocessed := processed

	md, err := r.Render(preprocessed)
	if err != nil {
		panic(err)
	}

	md = removeLinkStyleHack(md)

	var tagIdx = 0
	md = reTagBack.ReplaceAllStringFunc(md,
		func(s string) string {
			tag := tags[tagIdx]
			tagIdx++
			return TagsStyle.Render(tag)
		},
	)

	var lastLinkIdx = 0
	md = linkReplacementRE.ReplaceAllStringFunc(md,
		func(s string) string {
			linkID := s[2:5]
			active := docLinks.IsActive(tagIdx)
			sc := docLinks.GetShortcut(linkID)
			lastLinkIdx++

			return renderLink(linkID, sc, active, DocLinkStyles)
		},
	)

	return md, lastLinkIdx
}

func renderLink(link, sc string, active bool, styles LinkStyles) string {
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
