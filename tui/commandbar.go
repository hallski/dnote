package tui

import (
	"dnote/mdfiles"
	"dnote/tui/render"
	"strings"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const CommandBarHeight = 1

type commandBar struct {
	width, height int

	keymap commandBarKeymap

	textField textinput.Model

	statusMsg string
	statusId  int
}

func newCommandBar() commandBar {
	textfield := textinput.New()
	return commandBar{
		keymap:    defaultCmdKeyMap,
		textField: textfield,
		statusMsg: "",
		statusId:  1,
	}
}
func (m commandBar) Init() tea.Cmd {
	return nil
}

func (m commandBar) Update(msg tea.Msg) (commandBar, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		//
		case key.Matches(msg, m.keymap.Exit):
			m.textField.SetValue("")
			m.textField.Blur()
			return m, nil
		case key.Matches(msg, m.keymap.Commit):
			cmd := m.inputCmd()
			m.textField.SetValue("")
			m.textField.Blur()
			return m, cmd
		}

	case statusMsg:
		m.statusMsg = msg.s
		m.statusId++
		return m, timeoutStatusCmd(m.statusId)
	case statusMsgTimeoutMsg:
		if m.statusId == msg.id {
			m.statusMsg = ""
		}
		return m, nil
	}
	var cmd tea.Cmd
	m.textField, cmd = m.textField.Update(msg)
	return m, cmd
}

func (m commandBar) View() string {
	if !m.textField.Focused() {
		statusStyle := lipgloss.NewStyle().Width(m.width)
		return statusStyle.Render(m.statusMsg)
	}

	cmdStyle := lipgloss.NewStyle().Foreground(render.ColorHighCyan).Bold(true)
	normalStyle := lipgloss.NewStyle().Foreground(render.ColorWhite)

	styles := m.textField.Styles()
	if cmdEnd(m.textField.Value()) > 0 {
		styles.Focused.Text = cmdStyle
	} else {
		styles.Focused.Text = normalStyle
	}
	m.textField.SetStyles(styles)

	return m.textField.View()
}

type command struct {
	name    string
	hasArgs bool
	cmd     func(input string) tea.Cmd
}

var commands = []command{
	{
		"open",
		true,
		func(input string) tea.Cmd {
			return openLinkCmd(mdfiles.PadID(input))
		},
	},
	{
		"search",
		true,
		func(input string) tea.Cmd {
			return searchCmd(input)
		},
	},
	{
		"rand",
		false,
		func(input string) tea.Cmd {
			return emitMsgCmd(openRandomMsg{})
		},
	},
	{
		"last",
		false,
		func(input string) tea.Cmd {
			return emitMsgCmd(openLastMsg{})
		},
	},
	{
		// Reset collection
		"rc",
		false,
		func(input string) tea.Cmd {
			return emitMsgCmd(resetCollectionMsg{})
		},
	},
	{
		// Save to collection
		"sc",
		false,
		func(input string) tea.Cmd {
			return emitMsgCmd(saveToCollectionMsg{})
		},
	},
	{
		"add",
		true,
		func(input string) tea.Cmd {
			return addNoteCmd(input)
		},
	},
	{
		"inbox",
		true,
		func(input string) tea.Cmd {
			return emitMsgCmd(addInboxMsg{input})
		},
	},
}

func (m *commandBar) inputCmd() tea.Cmd {
	input := m.textField.Value()
	for _, c := range commands {
		prefix := c.name
		if c.hasArgs {
			prefix += " "
		}
		if strings.HasPrefix(input, prefix) {
			return c.cmd(input[len(prefix):])
		}
	}

	return emitMsgCmd(statusMsg{"Unknown command: " + input})
}

func cmdEnd(input string) int {
	for _, c := range commands {
		if strings.HasPrefix(input, c.name) {
			return len(c.name)
		}
	}

	return 0
}

func (m *commandBar) setSize(width, height int) {
	m.width, m.height = width, height
	m.textField.SetWidth(width)
}

func (m *commandBar) focus() tea.Cmd {
	return m.textField.Focus()
}

func (m *commandBar) inputActive() bool {
	return m.textField.Focused()
}

func (m *commandBar) startOpen(v string) tea.Cmd {
	m.textField.SetValue("open " + v)
	return m.textField.Focus()
}

func (m *commandBar) startAdd() tea.Cmd {
	cmd := "add"
	m.textField.SetValue(cmd + " ")
	return m.textField.Focus()
}

func (m *commandBar) startSearch(query string) tea.Cmd {
	input := "search " + query
	if query != "" {
		input += " "
	}

	m.textField.SetValue(input)
	return m.textField.Focus()
}

func (m *commandBar) startInbox() tea.Cmd {
	m.textField.SetValue("inbox ")
	return m.textField.Focus()
}
