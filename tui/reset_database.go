package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Update of the reset db page
func resetDatabaseUpdate(m home, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.page = 0
			return m, nil
		case "q", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

// View of the reset db page
func resetDatabaseView(m home) string {
	str := fmt.Sprintf("\n\n   %s Loading forever...press esc to quit\n\n", m.spinner.View())
	for _, frame := range m.spinner.Spinner.Frames {
		str += frame
	}
	return str
}
