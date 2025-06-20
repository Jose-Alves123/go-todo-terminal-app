// Package services implements utilities to support the main application,
// being with the use of os helper function or ainding with the arriging of
// data in the database
package services

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// GetAbsDatabasePath returns a string representing the absolute path to the database,
// given filename parameter of type string, that is the name of the database. For example,
// for the filename "hello.db" returnns absolute path ... /database/hello.db
func AbsDatabasePath(filename string) (string, error) {
	path, err := absProjectPath()

	if err != nil {
		return "", err
	}

	return path + "/database/" + filename, nil
}

// Gets full absolute path, including filename to the .env file.
//
// If root directory of project is called godoit, path wil be
// ... /godoit/.env
func EnvPath() (string, error) {
	path, err := absProjectPath()

	if err != nil {
		return "", err
	}

	return path + "/.env", nil
}

// Get filename of database, given information in the .env file
//
// The .env contains information about the name of the databases, for
// testing and development purposes
func EnvDatabaseFileName(fileLocation string, env string) (string, error) {
	err := godotenv.Load(fileLocation)
	if err != nil {
		return "Error loading .env file", err
	}

	return os.Getenv(env), nil
}

// Given parameter filePath as string, deletes file.
//
// document of filePath must be a file. If is directory
// or doucment does not exist, returns error.
func RemoveFile(filePath string) error {
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return errors.ErrUnsupported
	}

	return os.Remove(filePath)
}

// Get the full absolute path to the project root directory
func absProjectPath() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path[0 : strings.Index(path, "go-todo-terminal-app")+len("go-todo-terminal-app")], nil
}
