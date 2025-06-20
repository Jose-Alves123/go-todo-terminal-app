// Package main holds the main application
package main

import (
	"fmt"

	scripts "godoit/scripts"
	"godoit/services"
	"godoit/tui"
)

func main() {

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

	fmt.Println(path)

	if err != nil {
		panic("Automigration did not work")
	}

	tui.Start(db, path)
}
