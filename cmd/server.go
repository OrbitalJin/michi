package main

import (
	"log"
	"os"

	"github.com/OrbitalJin/michi/cli"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/server"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/store"
	"github.com/gin-gonic/gin"
)

var bangParserConfig = parser.NewConfig("!")
var shortcutParserConfig = parser.NewConfig("@")
var sesssionParserConfig = parser.NewConfig("#")

var storeConfig = store.NewConfig("./index.db")
var serviceConfig = service.NewConfig(true, "g")

var config = server.NewConfig(
	":5980",
	bangParserConfig,
	shortcutParserConfig,
	sesssionParserConfig,
	storeConfig,
	serviceConfig,
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	michi, err := server.New(config)

	if err != nil {
		panic(err)
	}

	michiCli := cli.New(michi)
	err = michiCli.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
