package database

import (
	"context"
	"database/sql"
	"github.com/erik-olsson-op/go-rest/internal/logger"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var Connection *sql.DB

// Init connection to the database server
func Init() {
	var err error
	Connection, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		logger.Logger.Fatal("Failed to connect to sqlite3 - api.db")
	}

	Connection.SetMaxOpenConns(10)
	Connection.SetMaxIdleConns(5)

	readScript("internal/database/scripts/ddl.sql")
	readScript("internal/database/scripts/dml.sql")
}

func readScript(path string) {
	script, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Fatalf("Failed to read script %s - %s", path, err)
	}

	_, err = Connection.ExecContext(context.Background(), string(script))
	if err != nil {
		logger.Logger.Fatal("Failed to exec script - ", err)
	}
	logger.Logger.Infof("Success - %s", path)
}
