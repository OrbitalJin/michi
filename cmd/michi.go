package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/OrbitalJin/michi/cli"
	"github.com/OrbitalJin/michi/internal"
	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/server"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if os.Getenv("ENV") == "dev" {
		gin.SetMode(gin.DebugMode)
		log.Println("Running in development mode.")
	}

	configDir, err := internal.EnsureConfigDir()
	if err != nil {
		log.Fatalf("Failed to prepare configuration directory: %v", err)
	}

	config, err := internal.LoadConfig(filepath.Join(configDir, "config.yaml"))
	if err != nil {
		log.Fatalf("Failed to load application configuration: %v", err)
	}

	config.Server.PidFile = internal.ExpandTilde(config.Server.PidFile)
	config.Server.LogFile = internal.ExpandTilde(config.Server.LogFile)
	config.Store.DBPath = internal.ExpandTilde(config.Store.DBPath)

	bangParserConfig := parser.NewConfig(config.Parser.BangPrefix)
	shortcutParserConfig := parser.NewConfig(config.Parser.ShortcutPrefix)
	sessionParserConfig := parser.NewConfig(config.Parser.SessionPrefix)

	storeConfig := store.NewConfig(config.Store.DBPath)
	serviceConfig := service.NewConfig(config.Service.KeepTrack, config.Service.DefaultProvider)

	serverConfig := server.NewConfig(
		config.Server.Port,
		config.Server.PidFile,
		config.Server.LogFile,
		bangParserConfig,
		shortcutParserConfig,
		sessionParserConfig,
		storeConfig,
		serviceConfig,
	)

	michi, err := server.New(serverConfig)

	if err != nil {
		panic(err)
	}

	michiCli := cli.New(michi)
	err = michiCli.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
