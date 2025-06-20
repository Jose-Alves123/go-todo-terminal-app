package tui

import (
	"godoit/models"
	"godoit/scripts"
	"godoit/services"
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

// Get the struct ToDoForm, with all necessary data to be used in the TUI
// for the form to create or edit a Task.
func createForm(title string, description string, state int) ToDoForm {
	titleInput := textinput.New()
	titleInput.SetValue(title)
	titleInput.Cursor.Style = CursorStyle
	titleInput.Width = 50
	titleInput.Placeholder = "Write the Title of the task here..."
	titleInput.Focus()
	titleInput.PromptStyle = FocusedStyle
	titleInput.TextStyle = FocusedStyle
	titleInput.CharLimit = 50

	descriptionInput := textarea.New()
	//descriptionInput.SetValue(description)
	descriptionInput.Cursor.Style = CursorStyle
	descriptionInput.Placeholder = "Write the Description of the task here..."
	descriptionInput.Blur()
	descriptionInput.CharLimit = 256

	stateOptions := stateButton{choices: []string{"To Do", "Doing", "Donne"}, selected: state}

	return ToDoForm{Title: titleInput, Description: descriptionInput, State: stateOptions, Len: 3}
}

// Manage teh form of the task to edit or create.
//
// Defines what element to be focused
func manageForm(m home) (home, tea.Cmd) {
	var cmd tea.Cmd
	switch m.focusIndex {
	case 0:
		cmd = m.inputs.Title.Focus()
		m.inputs.Title.PromptStyle = FocusedStyle
		m.inputs.Title.TextStyle = FocusedStyle
	case 1:
		cmd = m.inputs.Description.Focus()

	}
	return m, cmd
}

// Get tge table with the list of tasks.
//
// There are three columns: task, description and state
func getTable(m home) home {
	columns := []table.Column{
		{Title: "Task", Width: 10},
		{Title: "Description", Width: 30},
		{Title: "State", Width: 10},
	}

	rows := []table.Row{}
	toDos, _ := models.ToDos(m.db)
	for _, toDo := range toDos {
		rows = append(rows, table.Row{toDo.Title, toDo.Description, strconv.Itoa(int(toDo.State))})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m.table = t
	return m
}

// Create a spinner model
func generateSpinner() spinner.Model {
	s := spinner.New()
	s.Spinner = spinner.Moon
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))
	return s
}

// Reset database to be with empty data
func resetDB(m home) (*gorm.DB, string) {
	services.RemoveFile(m.path)
	rootPath, _ := services.EnvPath()
	dbFilename, _ := services.EnvDatabaseFileName(rootPath, "DEV_DB")
	path, _ := services.AbsDatabasePath(dbFilename)
	db, _ := scripts.AutoMigrate(path)
	return db, path
}
