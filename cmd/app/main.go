// Package main holds the main application
package main

import (
	"fmt"

	models "godoit/models"
	scripts "godoit/scripts"
	"godoit/services"
)

func main() {
	fmt.Println("Hello godoit")

	rootPath, err := services.EnvPath()
	if err != nil {
		panic("Error geting path to root of project")
	}

	dbFilename, err := services.EnvDatabaseFileName(rootPath, "DEV_DB")
	if err != nil {
		panic("Error loading .env file")
	}

	path, _ := services.AbsDatabasePath(dbFilename)
	db, err := scripts.AutoMigrate(path)

	if err != nil {
		panic("Automigration did not work")
	}

	title, description, state, err := services.DataToCreateToDo()

	if err != nil {
		panic("Erro a tentar obter dados")
	}
	fmt.Printf("Aqui estão os valores:\nTítulo: %s\nDescrição: %s\nEstado: %d\n", title, description, state)

	isValidated, err := models.ValidateCreation(&title, &description, &state)
	if err != nil || !isValidated {
		panic("Dados de todos não passaram na validação")
	}

	data, result := models.CreateToDo(title, description, state, db)

	if result.Error != nil {
		panic("An error occured data was not introduced in db")
	}

	fmt.Printf("New ToDo with id %d added to database\n", data.Id)

}
