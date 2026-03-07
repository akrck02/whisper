package tables

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func UpdateDatabaseTablesToLatestVersion(basePath string, db *sql.DB) error {

	latestDbVersion := 1

	// if no database exists, create one
	err := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='database_metadata'").Scan()
	if err == sql.ErrNoRows {
		return updateTablesForVersion(basePath, db, 0, latestDbVersion)
	}

	// get current version and update existing database
	var databaseVersion int
	err = db.QueryRow("SELECT version FROM database_metadata").Scan(&databaseVersion)
	if err != nil {
		return err
	}
	return updateTablesForVersion(basePath, db, databaseVersion, latestDbVersion)
}

func updateTablesForVersion(basePath string, db *sql.DB, currentVersion int, targetVersion int) error {

	for version := currentVersion + 1; version <= targetVersion; version++ {
		err := executeScriptIfExists(db, fmt.Sprintf("%s/sql/v%d/tables.sql", basePath, version))
		if nil != err {
			return err
		}

		err = executeScriptIfExists(db, fmt.Sprintf("%s/sql/v%d/data.sql", basePath, version))
		if nil != err {
			return err
		}

	}

	return nil
}

func executeScriptIfExists(db *sql.DB, fileName string) error {

	// if the file does not exist do not execute anything
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil
	}

	script, err := os.ReadFile(fileName)
	if nil != err {
		return err
	}

	statements := strings.SplitSeq(string(script), ";")

	for statement := range statements {
		_, err = db.Exec(statement)
		if nil != err {
			return err
		}
	}

	return nil
}
