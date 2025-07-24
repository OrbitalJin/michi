package main

import (
	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/server"
	"github.com/OrbitalJin/pow/internal/store"
)

var parserConfig = parser.NewConfig(`!(\b\w+\b)`, `!\b\w+\b`)
var storeConfig = store.NewConfig("./index.db", "g")
var config = server.NewConfig(parserConfig, storeConfig)

func main() {

	qmuxr, err := server.Default(config)
	if err != nil {
		panic(err)
	}

	qmuxr.Start()
}
