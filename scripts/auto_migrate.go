// Package scripts holds scripts that can be executed inside or outside
// of the main application
package scripts

import (
	"errors"
	"godoit/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Runs automigration of given database
func AutoMigrate(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return db, errors.ErrUnsupported
	}

	db.AutoMigrate(&models.ToDo{})
	return db, nil
}
