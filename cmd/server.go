package main

import (
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/server"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/OrbitalJin/qmuxr/internal/store"
)

var bangParserConfig = parser.NewConfig("!")
var shortcutParserConfig = parser.NewConfig("@")

var storeConfig = store.NewConfig("./index.db")
var serviceConfig = service.NewConfig(true, "g")

var config = server.NewConfig(bangParserConfig, shortcutParserConfig, storeConfig, serviceConfig)

func main() {

	qmuxr, err := server.Default(config)
	if err != nil {
		panic(err)
	}

	qmuxr.Start(":5980")
}
