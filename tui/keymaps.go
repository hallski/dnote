package tui

import "github.com/charmbracelet/bubbles/key"

type docKeymap struct {
	NextLink key.Binding
	PrevLink key.Binding
	OpenLink key.Binding
}

var DefaultAppKeyMap = appKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q", "ctrl+c"),
		key.WithHelp("ctrl+q", "quit"),
	),
	Search: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "search"),
	),
	AddNote: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add note"),
	),
	EditNode: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit note"),
	),
	Back: key.NewBinding(
		key.WithKeys("ctrl+o", "h"),
		key.WithHelp("ctrl+o or h", "back"),
	),
	Forward: key.NewBinding(
		key.WithKeys("tab", "l"),
		key.WithHelp("tab or l", "forward"),
	),
	StartCmd: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "start command"),
	),
	QuickOpen: key.NewBinding(
		key.WithKeys(getStrings(quickOpen)...),
		key.WithHelp("number", "start open"),
	),
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
		key.WithKeys("enter"),
		key.WithHelp("enter", "open link"),
	),
}
