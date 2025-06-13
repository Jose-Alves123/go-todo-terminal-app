// Package scripts_tes aims to test the functions of the package scripts
package scripts_test

import (
	"godoit/scripts"
	"godoit/services"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	envPath, _ := services.EnvPath()
	dbFilename, _ := services.EnvDatabaseFileName(envPath, "TEST_DB")
	path, _ := services.AbsDatabasePath(dbFilename)

	_, err := scripts.AutoMigrate(path)
	if err != nil {
		t.Errorf("Automigration should be successful")
	}
}
