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
		_, err := models.CreateToDo(tt.title, tt.description, tt.state, db)

		if (err != nil) != tt.wantError {
			t.Errorf(`Eror occured injecting data to database for Todo with parameters 
				Title - %s, 
				Description - %s, 
				State - %d, 
				expected error to be %t but was %t`,
				tt.title, tt.description, tt.state, tt.wantError, err != nil)
		}
	}

}

func TestEditToDo(t *testing.T) {
	var firstToDoUpdated models.ToDo
	var secondToDoUpdated models.ToDo

	envPath, _ := services.EnvPath()
	dbFilename, _ := services.EnvDatabaseFileName(envPath, "TEST_DB")
	path, _ := services.AbsDatabasePath(dbFilename)
	services.RemoveFile(path) // restart db data
	db, _ := scripts.AutoMigrate(path)

	firstToDo, _ := models.CreateToDo("My first task", "Description of my first task", uint8(0), db)
	secondToDo, _ := models.CreateToDo("My second task", "Description of my second task", uint8(1), db)

	err := models.EditToDo(firstToDo.ID, "My first task updated", "Description of my first task updated", uint8(1), db)
	db.First(&firstToDoUpdated, firstToDo.ID)

	if err != nil ||
		firstToDoUpdated.Title != "My first task updated" ||
		firstToDoUpdated.Description != "Description of my first task updated" ||
		firstToDoUpdated.State != uint8(1) {
		t.Errorf("Task with id %d does not match expect value", firstToDo.ID)
	}

	err = models.EditToDo(secondToDo.ID, "My second task updated", "Description of my second task updated", uint8(2), db)
	db.First(&secondToDoUpdated, secondToDo.ID)

	if err != nil ||
		secondToDoUpdated.Title != "My second task updated" ||
		secondToDoUpdated.Description != "Description of my second task updated" ||
		secondToDoUpdated.State != uint8(2) {
		t.Errorf("Task with id %d does not match expect value updated title - %s, description %s", secondToDoUpdated.ID, secondToDoUpdated.Title, secondToDoUpdated.Description)
	}

}

// Test reading of ToDo list
func TestToDos(t *testing.T) {
	envPath, _ := services.EnvPath()
	dbFilename, _ := services.EnvDatabaseFileName(envPath, "TEST_DB")
	path, _ := services.AbsDatabasePath(dbFilename)
	services.RemoveFile(path) // restart db data
	db, _ := scripts.AutoMigrate(path)

	result, err := models.ToDos(db)

	if err != nil || len(result) != 0 {
		t.Error("Databse should be empty in the beggining")
	}

	todo := models.ToDo{Title: "Task XPTO", Description: "Description of task XPTO", State: uint8(1)}
	db.Create(&todo)

	result, err = models.ToDos(db)
	if err != nil || len(result) != 1 {
		t.Error("Databse should be with one row at this point")
	}
}

// Test deletion of task from table todo
func TestDeleteToDo(t *testing.T) {
	envPath, _ := services.EnvPath()
	dbFilename, _ := services.EnvDatabaseFileName(envPath, "TEST_DB")
	path, _ := services.AbsDatabasePath(dbFilename)
	services.RemoveFile(path) // restart db data
	db, _ := scripts.AutoMigrate(path)

	firstToDo := models.ToDo{Title: "My first task", Description: "Description of my first task", State: uint8(0)}
	secondToDo := models.ToDo{Title: "My second task", Description: "Description of my second task", State: uint8(1)}
	db.Create(&firstToDo)
	db.Create(&secondToDo)

	twoToDos, err := models.ToDos(db)
	if err != nil || len(twoToDos) != 2 {
		t.Error("Databse should be with two rows at this point")
	}

	models.DeleteToDo(firstToDo, db)

	oneToDo, err := models.ToDos(db)
	if err != nil || len(oneToDo) != 1 {
		t.Error("Databse should be with one row at this point")
	}

	models.DeleteToDo(secondToDo, db)

	zeroToDo, err := models.ToDos(db)
	if err != nil || len(zeroToDo) != 0 {
		t.Error("Databse should be with no rows at this point")
	}
}
