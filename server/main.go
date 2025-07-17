package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB = "./foo.db"

func main() {
	os.Remove(DB)
	db, err := sql.Open("sqlite3", DB)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`

	_, err = db.Exec(sqlStmt)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
