package tui

import (
	"dnote/mdfiles"
	"dnote/render"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	textfield.Cursor.Blink = true
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
	case tea.KeyMsg:
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
	style := lipgloss.NewStyle().Foreground(render.ColorWhite)

	if cmdEnd(m.textField.Value()) > 0 {
		m.textField.TextStyle = cmdStyle
	} else {
		m.textField.TextStyle = style
	}

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
			return addNoteCmd(input, false)
		},
	},
	{
		"Add",
		true,
		func(input string) tea.Cmd {
			return addNoteCmd(input, true)
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
	m.textField.Width = width
}

func (m *commandBar) focus() tea.Cmd {
	return m.textField.Focus()
}

func (m *commandBar) focused() bool {
	return m.textField.Focused()
}

func (m *commandBar) startOpen(v string) tea.Cmd {
	m.textField.SetValue("open " + v)
	return m.textField.Focus()
}

func (m *commandBar) startAdd(keepFocus bool) tea.Cmd {
	cmd := "add"
	if keepFocus {
		cmd = "Add"
	}
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
