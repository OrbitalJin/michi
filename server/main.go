package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

const DB = "./foo.db"

func main() {

	pattern := `!\b\w+\b`
	target := "!t3 fdsfdslj !213 !go ! hello"
	re, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal(err)
	}

	matches := re.FindAllString(target, -1)

	if len(matches) > 0 {
		for i, match := range matches {
			fmt.Printf("match %d, %s \n", i, match)
		}
	}

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
