package tui

import "github.com/charmbracelet/bubbles/key"

type appKeyMap struct {
	Quit           key.Binding
	Search         key.Binding
	EditNode       key.Binding
	Back           key.Binding
	Forward        key.Binding
	StartCmd       key.Binding
	QuickOpen      key.Binding
	AddNote        key.Binding
	OpenRandomNote key.Binding
	OpenLastNote   key.Binding
}

type searchKeymap struct {
	NextLink     key.Binding
	PrevLink     key.Binding
	OpenLink     key.Binding
	ExtendSearch key.Binding
}

type docKeymap struct {
	NextLink key.Binding
	PrevLink key.Binding
	OpenLink key.Binding
}

type commandBarKeymap struct {
	Exit   key.Binding
	Commit key.Binding
}

var quickOpen = []byte("0123456789")

func getStrings(bytes []byte) []string {
	var ss []string

	for _, b := range bytes {
		ss = append(ss, string(b))
	}
	return ss
}

var defaultAppKeyMap = appKeyMap{
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q", "q"),
		key.WithHelp("ctrl+q", "quit"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	),
	EditNode: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit note"),
	),
	Back: key.NewBinding(
		key.WithKeys("ctrl+o"),
		key.WithHelp("ctrl+o", "back"),
	),
	Forward: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "forward"),
	),
	StartCmd: key.NewBinding(
		key.WithKeys("."),
		key.WithHelp(".", "start command"),
	),
	QuickOpen: key.NewBinding(
		key.WithKeys(getStrings(quickOpen)...),
		key.WithHelp("number", "start open"),
	),
	AddNote: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add note"),
	),
	OpenRandomNote: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "randoM note"),
	),
	OpenLastNote: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "last note"),
	),
}

var defaultDocKeyMap = docKeymap{
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

var defaultSearchKeyMap = searchKeymap{
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
	ExtendSearch: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "edit search"),
	),
}

var defaultCmdKeyMap = commandBarKeymap{
	Exit: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "exit command"),
	),
	Commit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "commit command"),
	),
}
