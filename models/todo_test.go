// Package models_test aims to test the function of the package models
package models_test

import (
	"godoit/models"
	"godoit/scripts"
	"godoit/services"
	"testing"
)

// Tests different cases of validation of ToDo
func TestValidateCreation(t *testing.T) {
	data := []struct {
		title       string
		description string
		state       uint8
		wantError   bool
	}{
		{title: "Tarefa 1", description: "Descrição da tarefa 1", state: 1, wantError: false},
		{title: "T2", description: "Descrição da tarefa 2", state: 1, wantError: true},
		{title: "Tarefa 3", description: "Descrição da tarefa 3", state: 4, wantError: true},
		{title: "Tarefa 4", description: "Descrição da tarefa 4, Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,", state: 4, wantError: true},
		{title: "Tarefa já completa", description: "Descrição da tarefa já completa", state: 2, wantError: false},
		{title: "Tarefa por fazer", description: "Descrição da tarefa para fazer", state: 0, wantError: false},
	}

	for _, tt := range data {
		_, err := models.ValidateCreation(&tt.title, &tt.description, &tt.state)

		if (err != nil) != tt.wantError {
			t.Errorf(`ValidateCreation("%s", "%s", %d) , won't match validation. expected error to be %t`, tt.title, tt.description, tt.state, tt.wantError)
		}
	}

}

// Tests different cases of addition of ToDo to database
func TestCreateToDo(t *testing.T) {
	envPath, _ := services.EnvPath()
	dbFilename, _ := services.EnvDatabaseFileName(envPath, "TEST_DB")
	path, _ := services.AbsDatabasePath(dbFilename)
	db, _ := scripts.AutoMigrate(path)
	data := []struct {
		title       string
		description string
		state       uint8
		wantError   bool
	}{
		{title: "Tarefa 1", description: "Descrição da tarefa 1", state: 1, wantError: false},
		{title: "T2", description: "Descrição da tarefa 2", state: 1, wantError: true},
		{title: "Tarefa 3", description: "Descrição da tarefa 3", state: 4, wantError: true},
		{title: "Tarefa 4", description: "Descrição da tarefa 4, Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,Descrição da tarefa 4,", state: 4, wantError: true},
		{title: "Tarefa já completa", description: "Descrição da tarefa já completa", state: 2, wantError: false},
		{title: "Tarefa por fazer", description: "Descrição da tarefa para fazer", state: 0, wantError: false},
	}

	for _, tt := range data {
		_, result := models.CreateToDo(tt.title, tt.description, tt.state, db)

		if (result.Error != nil) != tt.wantError {
			t.Errorf(`Eror occured injecting data to database for Todo with parameters 
				Title - %s, 
				Description - %s, 
				State - %d, 
				expected error to be %t but was %t`,
				tt.title, tt.description, tt.state, tt.wantError, result.Error != nil)
		}
	}

}
