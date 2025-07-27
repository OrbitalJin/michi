package server

import (
	"fmt"

	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/server/handler"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/OrbitalJin/qmuxr/internal/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store           *store.Store
	router          *gin.Engine
	handler         *handler.Handler
	queryParser     *parser.QueryParser
	providerService *service.SearchProviderService
	historyService  *service.HistoryService
	config          *Config
}

func New(config *Config, useCors bool) (*Server, error) {
	store, err := store.New(config.storeCfg)

	if err != nil {
		return nil, fmt.Errorf("failed to access the store: %w", err)
	}

	qp, err := parser.NewQueryParser(config.bangParserCfg, config.shortcutParserCfg)

	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	psvc := service.NewSearchProviderService(
		qp.BangParser(),
		store.SearchProviders,
		config.serviceCgf,
	)

	hsvc := service.NewHistoryService(store.History)

	handler := handler.NewHandler(qp, psvc, hsvc, "q")

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
		providerService: psvc,
		historyService:  hsvc,
		queryParser:     qp,
		router:          router,
		handler:         handler,
		store:           store,
		config:          config,
	}, nil
}

func Default(config *Config) (*Server, error) {
	server, err := New(config, false)

	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}

	if err = server.store.Migrate(); err != nil {
		return nil, fmt.Errorf("failed to conduct database migration: %w", err)
	}

	server.router.GET("/search", func(ctx *gin.Context) {
		server.handler.Root(ctx)
	})

	return server, nil
}

func (sv *Server) Start() {
	sv.router.Run()
}
