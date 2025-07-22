package main

import (
	"github.com/OrbitalJin/pow/internal/store"
)

func main() {
	db, err := store.Init("./data/database.db")

	if err != nil {
		panic(err)
	}

	err = db.Migrate()

	if err != nil {
		panic(err)
	}

	err = store.Import("./data/search_providers.json", db)

	if err != nil {
		panic(err)
	}
}
