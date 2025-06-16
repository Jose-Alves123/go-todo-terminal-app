// Package services_test aims to test the functions of the package services
package services_test

import (
	"godoit/scripts"
	"godoit/services"
	"testing"
)

// Test getting the absolute database path
func TestAbsDatabasePath(t *testing.T) {
	_, err := services.AbsDatabasePath("cooldatabase.db")

	if err != nil {
		t.Errorf("Failed to obtain the full database path")
	}
}

// Tests get absolute .env path
func TestEnvPath(t *testing.T) {
	_, err := services.EnvPath()

	if err != nil {
		t.Errorf("Failed to obtain the full .env path")
	}
}

// Tests get filename of database from .env file
//
// There is a database for development and a database for testing.
// This function aims to check if we're getting the correct
// database filename
func TestEnvDatabaseFileName(t *testing.T) {
	rootPath, _ := services.EnvPath()

	envs := []struct {
		title        string
		wantError    bool
		isEmptyValue bool
	}{
		{title: "DEV_DB", wantError: false, isEmptyValue: false},
		{title: "TEST_DB", wantError: false, isEmptyValue: false},
		{title: "UMPA_LUMPA", wantError: false, isEmptyValue: true},
	}

	for _, tt := range envs {
		dbFilename, err := services.EnvDatabaseFileName(rootPath, tt.title)

		if (err != nil) != tt.wantError || (len(dbFilename) == 0) != tt.isEmptyValue {
			t.Errorf("Failed for title - %s %T", tt.title, err)
		}
	}

}

// Tests remove a file
//
// A file is removed if exists and is indeed a file. Cannot be a directory
func TestRemoveFile(t *testing.T) {
	envPath, _ := services.EnvPath()
	dbFilename, _ := services.EnvDatabaseFileName(envPath, "TEST_DB")
	path, _ := services.AbsDatabasePath(dbFilename)
	scripts.AutoMigrate(path)

	err := services.RemoveFile(path)

	if err != nil {
		t.Errorf("Should have deleted test db successfully, but didn't")
	}

	err = services.RemoveFile(path)

	if err == nil {
		t.Errorf("Should have thrown an error, because db should not exist at this point")
	}
}
