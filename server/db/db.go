package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db   *sql.DB
	path string
}

func Init(path string) (error, *DB) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return err, nil
	}
	return nil, &DB{
		db, path,
	}
}
