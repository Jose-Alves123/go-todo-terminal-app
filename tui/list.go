package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// View of the listing of tasks
func listView(m home) string {
	return m.table.View() + "\n\nPres esc to go back\n\n" + fmt.Sprintf("%d %d", m.cursor, len(m.table.Rows()))
}

// Update of the listing of tasks page
func listUpdate(msg tea.Msg, m home) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.page = 0
			return m, nil
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.table.Rows())-1 {
				m.cursor++
			}
		case "backspace":
			m.page = 0
		}

	}

	m.table.SetCursor(m.cursor)

	return m, nil
}
