package tui

import (
	"dnote/mdfiles"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type commandBar struct {
	input string

	width int

	keymap commandBarKeymap
}

func (cb commandBar) Init() tea.Cmd {
	return nil
}

func (cb commandBar) Update(msg tea.Msg) (commandBar, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, cb.keymap.Exit):
			cb.input = ""
			return cb, exitCommandBar
		case key.Matches(msg, cb.keymap.Commit):
			cmd := cb.inputCmd()
			cb.input = ""
			return cb, tea.Batch(cmd, exitCommandBar)
		case key.Matches(msg, cb.keymap.Backspace):
			cb.backspace()
		default:
			cb.input += msg.String()
		}
	}

	return cb, nil
}

func (cb commandBar) View() string {
	promptStyle := lipgloss.NewStyle().Foreground(colorHighBlue).Bold(true).MarginRight(1)

	cmdStyle := lipgloss.NewStyle().Foreground(colorHighCyan).Bold(true)

	style := lipgloss.NewStyle().Foreground(colorWhite)

	cmdLen := cmdEnd(cb.input)
	cmd := cmdStyle.Width(cmdLen).Render(cb.input[:cmdLen])
	rest := style.Width(cb.width - cmdLen).Render(cb.input[cmdLen:])

	prompt := promptStyle.Render("❯")

	return prompt + cmd + rest
}

func newCommandBar() commandBar {
	return commandBar{
		"",
		0,
		defaultCmdKeyMap,
	}
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
			return openLinkCmd(mdfiles.PadID(input[5:]))
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
}

func (cb *commandBar) inputCmd() tea.Cmd {
	for _, c := range commands {
		prefix := c.name
		if c.hasArgs {
			prefix += " "
		}
		if strings.HasPrefix(cb.input, prefix) {
			return c.cmd(cb.input)
		}
	}

	return emitMsgCmd(statusMsg{"Unknown command: " + cb.input})
}

func cmdEnd(input string) int {
	for _, c := range commands {
		if strings.HasPrefix(input, c.name) {
			return len(c.name)
		}
	}

	return 0
}

func (cb *commandBar) backspace() {
	length := len(cb.input)
	if length <= 0 {
		return
	}

	cb.input = cb.input[:length-1]
}

func (cb *commandBar) setWidth(width int) {
	cb.width = width
}

func (cb *commandBar) startOpen(v string) {
	cb.input = "open " + v
}
