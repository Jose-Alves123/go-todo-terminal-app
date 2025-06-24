package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// View of the listing of tasks
func listView(m home) string {
	return m.table.View() + "\n\nPres esc to go back\n\n" + fmt.Sprintf("press e to edit, d to delete, esc to return back, ctrl+c to exit\n%d %d", m.cursor, len(m.table.Rows()))
}

// Update of the listing of tasks page
func listUpdate(msg tea.Msg, m home) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit
		case "d":
			m = deleteToDo(m)
		case "e":
			m.page = 2
			m.inputs = editForm(m.todos[m.cursor].Title, m.todos[m.cursor].Description, int(m.todos[m.cursor].State), m.todos[m.cursor].ID)
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
