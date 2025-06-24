// Package models holds the database models with its diferent functions
package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// A ToDo holds information about a task that may or
// may not have been completed
type ToDo struct {
	gorm.Model
	ID          int            `gorm:"index;primaryKey;autoIncrement:true;unique"`
	Title       string         `gorm:"column:title;not null;check:(length(title)>=5 and length(title)<=50)"`
	Description string         `gorm:"column:description;type:text;check:length(description)<=256"`
	State       uint8          `gorm:"column:state;not null;default:0;check:state>=0 and state<=2"`
	CreatedAt   time.Time      `gorm:"index;not null;default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"index;not null;default:current_timestamp"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Validates the title, description and state of a todo.
//
// Check ToDo struct to see all the validation principles
func ValidateCreation(title *string, description *string, state *uint8) (bool, error) {
	passesValidation := title != nil && len(*title) >= 5 && len(*title) <= 50 &&
		(description == nil || len(*description) <= 265) &&
		state != nil && *state <= 2
	if !passesValidation {
		return passesValidation, errors.ErrUnsupported
	}
	return passesValidation, nil

}

// Adds ToDo to dabase
func CreateToDo(title string, description string, state uint8, db *gorm.DB) (ToDo, error) {
	toDo := ToDo{Title: title, Description: description, State: uint8(state)}
	result := db.Create(&toDo)
	return toDo, result.Error
}

// Èdits task with giben ID
//
// Returns error if any
func EditToDo(id int, title string, description string, state uint8, db *gorm.DB) error {
	result := db.Model(&ToDo{ID: id}).Updates(ToDo{
		Title:       title,
		Description: description,
		State:       state,
	})

	return result.Error
}

// Delete task from ToDo table
func DeleteToDo(toDo ToDo, db *gorm.DB) {
	db.Delete(&toDo)
}

// Get all ToDos from database
func ToDos(db *gorm.DB) ([]ToDo, error) {
	var toDos []ToDo
	result := db.Find(&toDos)
	return toDos, result.Error
}
