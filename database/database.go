package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/akrck02/whisper/sdk/logger"
)

func GetConnection() (*sql.DB, error) {
	db, err := Connect("whisper.db")
	if nil != err {
		logger.Error(err)
		return nil, err
	}

	db.SetMaxOpenConns(100)
	return db, nil
}

func Close(c *sql.DB) {
	if nil == c {
		return
	}
	c.Close()
}

func Connect(filePath string) (*sql.DB, error) {
	path := fmt.Sprintf("%s?cache=shared&mode=rwc&_journal=WAL&_foreign_keys=on", filePath)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}
