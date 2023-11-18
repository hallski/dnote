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
}

func newCommandBar() commandBar {
	textfield := textinput.New()
	textfield.Cursor.Blink = true
	return commandBar{
		keymap:    defaultCmdKeyMap,
		textField: textfield,
	}
}
func (cb commandBar) Init() tea.Cmd {
	return nil
}

func (cb commandBar) Update(msg tea.Msg) (commandBar, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		//
		case key.Matches(msg, cb.keymap.Exit):
			cb.textField.SetValue("")
			cb.textField.Blur()
			return cb, nil
		case key.Matches(msg, cb.keymap.Commit):
			cmd := cb.inputCmd()
			cb.textField.SetValue("")
			cb.textField.Blur()
			return cb, cmd
		}

	}
	var cmd tea.Cmd
	cb.textField, cmd = cb.textField.Update(msg)
	return cb, cmd
}

func (cb commandBar) View() string {

	cmdStyle := lipgloss.NewStyle().Foreground(render.ColorHighCyan).Bold(true)
	style := lipgloss.NewStyle().Foreground(render.ColorWhite)

	if cmdEnd(cb.textField.Value()) > 0 {
		cb.textField.TextStyle = cmdStyle
	} else {
		cb.textField.TextStyle = style
	}

	return cb.textField.View()
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

func (cb *commandBar) inputCmd() tea.Cmd {
	input := cb.textField.Value()
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

func (cb *commandBar) setSize(width, height int) {
	cb.width, cb.height = width, height
	cb.textField.Width = width
}

func (cb *commandBar) focus() tea.Cmd {
	return cb.textField.Focus()
}

func (cb *commandBar) blur() {
	cb.textField.Blur()
}

func (cb *commandBar) focused() bool {
	return cb.textField.Focused()
}

func (cb *commandBar) startOpen(v string) tea.Cmd {
	cb.textField.SetValue("open " + v)
	return cb.textField.Focus()
}

func (cb *commandBar) startAdd(keepFocus bool) tea.Cmd {
	cmd := "add"
	if keepFocus {
		cmd = "Add"
	}
	cb.textField.SetValue(cmd + " ")
	return cb.textField.Focus()
}

func (cb *commandBar) startSearch(query string) tea.Cmd {
	input := "search " + query
	if query != "" {
		input += " "
	}

	cb.textField.SetValue(input)
	return cb.textField.Focus()
}
