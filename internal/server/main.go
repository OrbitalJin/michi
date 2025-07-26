package server

import (
	"fmt"
	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/OrbitalJin/qmuxr/internal/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router          *gin.Engine
	ProviderService *service.ProviderService
	HistoryService  *service.HistoryService
	parser          *parser.Parser
	store           *store.Store
	config          *Config
}

func New(config *Config, useCors bool) (*Server, error) {
	parser, err := parser.NewParser(config.parserCfg)

	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	store, err := store.New(config.storeCfg)

	if err != nil {
		return nil, fmt.Errorf("failed to access the store: %w", err)
	}

	pService := service.NewProviderService(
		parser,
		store.SearchProviders,
		config.serviceCgf,
	)

	hService := service.NewHistoryService(store.History)

	router := gin.Default()

	if useCors {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
	}

	return &Server{
		Router:          router,
		ProviderService: pService,
		HistoryService:  hService,
		parser:          parser,
		store:           store,
		config:          config,
	}, nil
}

func Default(config *Config) (*Server, error) {
	server, err := New(config, true)

	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}

	if err = server.store.Migrate(); err != nil {
		return nil, fmt.Errorf("failed to conduct database migration: %w", err)
	}

	server.Router.GET("/search", func(ctx *gin.Context) {
		Search(
			ctx,
			server.ProviderService,
			server.HistoryService,
		)
	})

	return server, nil
}

func (sv *Server) Start() {
	sv.Router.Run()
}
