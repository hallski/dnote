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
	cmdStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#aa00aa")).
		Foreground(lipgloss.Color("#ffff55"))

	style := lipgloss.NewStyle().
		Background(lipgloss.Color("#0000aa")).
		Foreground(lipgloss.Color("#ffff55")).
		Width(cb.width)

	cmdLen := cmdEnd(cb.input)
	return cmdStyle.Render(cb.input[:cmdLen]) + style.Render(cb.input[cmdLen:])

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
		prefix := c.name
		if c.hasArgs {
			prefix += " "
		}
		if strings.HasPrefix(input, prefix) {
			return len(prefix)
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
