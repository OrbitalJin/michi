package app

import (
	"fmt"

	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/store"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Parser *parser.Parser
	Store *store.DB
	cfg *Config
}

func New(appCfg *Config) *App {
	parser, err := parser.NewParser(appCfg.parserCfg)

	if err != nil {
		panic(fmt.Errorf("failed to start the app, couldn't get a parser: %w", err))
	}

	store, err := store.New(appCfg.storeCfg)

	if err != nil {
		panic(fmt.Errorf("failed to initiate the connection with the database: %w", err))
	}

	return &App{
		Router: gin.Default(),
		Parser: parser,
		Store: store,
		cfg: appCfg,
	}
}
