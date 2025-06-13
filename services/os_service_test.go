// Package services_test aims to test the functions of the package services
package services_test

import (
	"godoit/services"
	"strings"
	"testing"
)

func TestGetAbsDatabasePath(t *testing.T) {
	path, err := services.AbsDatabasePath("cooldatabase.db")
	suffix := "go-todo-terminal-app/database/cooldatabase.db"

	if err != nil || !strings.HasSuffix(path, suffix) {
		t.Errorf("Database path does not match expected %s to have suffix %s", path, suffix)
	}
}

func TestAbsDatabasePath(t *testing.T) {
	_, err := services.AbsDatabasePath("cooldatabase.db")

	if err != nil {
		t.Errorf("Failed to obtain the full database path")
	}
}

func TestEnvPath(t *testing.T) {
	_, err := services.EnvPath()

	if err != nil {
		t.Errorf("Failed to obtain the full .env path")
	}
}

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
