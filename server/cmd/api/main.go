package main

import (
	"github.com/OrbitalJin/pow/internal/database"
)

func main() {
	db, err := database.Init("./foo.db")

	if err != nil {
		panic(err)
	}

	err = db.Migrate()

	if err != nil {
		panic(err)
	}
}
