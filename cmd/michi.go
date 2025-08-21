package main

import (
	"fmt"
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
	isDebug := os.Getenv("ENV") == "dev"

	if isDebug {
		gin.SetMode(gin.DebugMode)
		log.Println("Running in development mode.")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	configDir, err := internal.EnsureConfigDir()
	if err != nil {
		log.Fatalf("Failed to prepare configuration directory: %v", err)
	}

	configPath := filepath.Join(configDir, "config.yaml")
	config, err := internal.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load application configuration: %v", err)
	}

	if err = internal.EnsureHydrationFile(); err != nil {
		log.Fatalf("Failed to hydrate database: %v", err)
	}

	bangParserConfig := parser.NewConfig(config.Parser.BangPrefix)
	shortcutParserConfig := parser.NewConfig(config.Parser.ShortcutPrefix)
	sessionParserConfig := parser.NewConfig(config.Parser.SessionPrefix)

	storeConfig := store.NewConfig(config.DBPath)
	serviceConfig := service.NewConfig(config.Service.KeepTrack, config.Service.DefaultProvider)

	serverConfig := server.NewConfig(
		config.Server.Port,
		config.PidFile,
		config.LogFile,
		bangParserConfig,
		shortcutParserConfig,
		sessionParserConfig,
		storeConfig,
		serviceConfig,
	)

	michi, err := server.New(serverConfig)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	michiCli := cli.New(michi)
	if err := michiCli.Run(os.Args); err != nil {
		fmt.Printf("%s●%s %v\n", internal.ColorRed, internal.ColorReset, err)
	}
}
