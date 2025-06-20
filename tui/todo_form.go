package tui

import (
	"fmt"
	"godoit/models"

	tea "github.com/charmbracelet/bubbletea"
)

// Update of the form to create/edit a task
func formUpdate(msg tea.Msg, m home) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.page = 0
			return m, nil
		case "enter":
			if m.focusIndex == (m.inputs.Len + len(m.inputs.State.choices) - 1) {
				titleVal := m.inputs.Title.Value()
				descriptionVal := m.inputs.Description.Value()
				stateVal := uint8(m.inputs.State.selected)
				passes, _ := models.ValidateCreation(&titleVal, &descriptionVal, &stateVal)
				if passes {
					models.CreateToDo(titleVal, descriptionVal, stateVal, m.db)
					m.page = 0
				}
				return m, nil
			} else if m.focusIndex > m.inputs.Len-2 {
				m.inputs.State.selected = m.focusIndex - m.inputs.Len + 1
			}
			m, cmd = manageForm(m)
			return m, tea.Batch(cmd)

		case "tab", "shift+tab", "up", "down":
			s := msg.String()

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > m.inputs.Len+len(m.inputs.State.choices)-1 {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = m.inputs.Len + len(m.inputs.State.choices) - 1
			}

			m.inputs.Title.Blur()
			m.inputs.Title.PromptStyle = NoStyle
			m.inputs.Title.TextStyle = NoStyle

			m.inputs.Description.Blur()

			m, cmd = manageForm(m)
			return m, tea.Batch(cmd)
		}
	}

	cmd = m.updateInputs(msg)
	return m, cmd
}

// Update inputs
func (m *home) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, m.inputs.Len)
	m.inputs.Title, cmds[0] = m.inputs.Title.Update(msg)
	m.inputs.Description, cmds[1] = m.inputs.Description.Update(msg)
	return tea.Batch(cmds...)
}

// View of the form to create/edit a task
func formView(m home) string {
	b := "Title\n"
	b += fmt.Sprintf("%s\n", m.inputs.Title.View())

	b += "\n\ndescription\n"
	b += fmt.Sprintf("%s\n\n", m.inputs.Description.View())

	for i, option := range m.inputs.State.choices {
		cursor := " "
		if m.focusIndex == m.inputs.Len+i-1 {
			cursor = ">"
		}
		mark := "   "
		if m.inputs.State.selected == i {
			mark = " X "
		}
		b += fmt.Sprintf("%s [%s] %s\n", cursor, mark, option)
	}

	button := &BlurredButton
	if m.focusIndex == m.inputs.Len+len(m.inputs.State.choices)-1 {
		button = &FocusedButton
	}
	b += fmt.Sprintf("\n\n%s\n\n", *button)

	return b
}
