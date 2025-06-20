package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// View in the home page
func homeView(m home) string {
	s := "GoDoIt\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s  %s\n", cursor, choice)
	}

	s += "\nPress esc to quit.\n"
	return s
}

// Update in the home page
func homeUpdate(msg tea.Msg, m home) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.page = m.cursor + 1
			m.cursor = 0

			if m.page == 1 {
				m = getTable(m)
			}
			if m.page == 2 {
				m.inputs = createForm("", "", 0)
			}

			if m.page == 3 {
				m.db, m.path = resetDB(m)
				m.spinner = generateSpinner()
				cmd = m.spinner.Tick
			}

		}
	}

	return m, cmd
}
