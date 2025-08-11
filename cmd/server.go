package main

import (
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/server"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/store"
)

var bangParserConfig = parser.NewConfig("!")
var shortcutParserConfig = parser.NewConfig("@")
var sesssionParserConfig = parser.NewConfig("#")

var storeConfig = store.NewConfig("./index.db")
var serviceConfig = service.NewConfig(true, "g")

var config = server.NewConfig(
	bangParserConfig,
	shortcutParserConfig,
	sesssionParserConfig,
	storeConfig,
	serviceConfig,
)

func main() {

	michi, err := server.New(config)
	if err != nil {
		panic(err)
	}

	michi.Start(":5980")
}
