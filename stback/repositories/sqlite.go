package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDb struct {
	db *sql.DB
}

func newSqliteDb(dbPath string) (*SqliteDb, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &SqliteDb{db: db}, nil
}

func (sdb *SqliteDb) close() {
	if sdb.db != nil {
		sdb.db.Close()
	}
}
