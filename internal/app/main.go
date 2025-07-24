package app

import (
	"fmt"

	"github.com/OrbitalJin/pow/internal/parser"
	"github.com/OrbitalJin/pow/internal/service"
	"github.com/OrbitalJin/pow/internal/store"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router  *gin.Engine
	Service *service.ProviderService
	parser  *parser.Parser
	store   *store.Store
	cfg     *Config
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

	service := service.NewProviderService(parser, store)

	return &App{
		Router:  gin.Default(),
		Service: service,
		parser:  parser,
		store:   store,
		cfg:     appCfg,
	}
}

func (app *App) Start() {
	app.Router.Run()
}
