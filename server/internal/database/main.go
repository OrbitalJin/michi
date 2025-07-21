package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
	path string
}

func Init(path string) (*DB, error) {
	if path == "" {
		return nil, fmt.Errorf("Path to database cannot be nil.")
	}

	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}

	return &DB{
		conn: db,
		path: path,
	}, nil
}

func (db *DB) Migrate() error {
	stmnt := `
	CREATE TABLE IF NOT EXISTS bangs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);`

	_, err := db.conn.Exec(stmnt)

	if err != nil {
		return err
	}

	return nil
}
