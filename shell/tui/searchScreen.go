package tui

import (
	"fmt"
	"hallski/thinkadus/core"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item core.Note

// Implement the list.DefaultItem interface for Note
func (i item) Title() string {
	return i.Name
}

func (i item) Description() string {
	return fmt.Sprintf("%v bytes", i.Size)
}
func (i item) FilterValue() string {
	return i.Name
}

type ListNotesModel struct {
	list list.Model
}

func (m ListNotesModel) Init() tea.Cmd {
	return nil
}

func (m ListNotesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Printf("Window size changed to %vx%v\n", msg.Width, msg.Height)
		m.list.SetHeight(msg.Height)
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.list.SettingFilter() {
				var cmd tea.Cmd
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			}
			if i, ok := m.list.SelectedItem().(item); ok {
				return m, NavigateToNote(core.Note(i))
			}
			// Open selected note
			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		default:
			var cmd tea.Cmd
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}
	default:
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
}

func (m ListNotesModel) View() string {
	return m.list.View()
}

func makeListItems(notes []core.Note) []list.Item {
	items := make([]list.Item, len(notes))
	for i, note := range notes {
		items[i] = item(note)
	}
	return items
}

func InitListNotesScreen(notes []core.Note, width, height int) tea.Model {
	var l = list.New(makeListItems(notes), list.NewDefaultDelegate(), width, height)
	l.Title = "Notes"
	l.SetShowFilter(true)

	return ListNotesModel{list: l}
}
