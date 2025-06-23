package tui

import (
	"fmt"
	"godoit/models"
	"os"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

var (
	FocusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).MarginRight(1)
	BlurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginRight(1)
	CursorStyle   = FocusedStyle
	NoStyle       = lipgloss.NewStyle()
	FocusedButton = FocusedStyle.Render("[ Submit ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Submit"))
)

type stateButton struct {
	choices  []string
	selected int
}

type ToDoForm struct {
	Title       textinput.Model
	Description textarea.Model
	State       stateButton
	Len         int
}

type home struct {
	tea.Model
	choices    []string
	todos      []models.ToDo
	table      table.Model
	cursor     int
	formCursor cursor.Mode
	focusIndex int
	inputs     ToDoForm
	page       int
	db         *gorm.DB
	path       string
	spinner    spinner.Model
}

// Start of the TUI aplication, in full-screen mode
func Start(db *gorm.DB, path string) {
	p := tea.NewProgram(model(db, path), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// Get the initial model
func model(db *gorm.DB, path string) home {

	m := home{
		choices:    []string{"List ToDos", "Create ToDo", "Restart Database"},
		page:       0,
		cursor:     0,
		formCursor: cursor.CursorBlink,
		inputs:     ToDoForm{},
		db:         db,
		path:       path,
	}

	return m
}

// Initializion of the programm with defined commands
func (m home) Init() tea.Cmd {
	return tea.Batch(tea.SetWindowTitle("GoDoIt"), textinput.Blink, textarea.Blink)
}

// Update of the program
func (m home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.page {
	case 0:
		return homeUpdate(msg, m)
	case 1:
		return listUpdate(msg, m)
	case 2:
		return formUpdate(msg, m)
	case 3:
		return resetDatabaseUpdate(m, msg)
	default:
		return homeUpdate(msg, m)
	}

}

// View of the program
func (m home) View() string {
	switch m.page {
	case 0: // homepage
		return homeView(m)
	case 1: // list page
		return listView(m)
	case 2: // create or edit todo
		return formView(m)
	case 3: // reseting database loading page
		return resetDatabaseView(m)
	default:
		return homeView(m)
	}

}
