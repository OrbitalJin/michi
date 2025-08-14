package server

import (
	"fmt"

	"github.com/OrbitalJin/michi/internal/parser"
	"github.com/OrbitalJin/michi/internal/router"
	"github.com/OrbitalJin/michi/internal/router/handler"
	"github.com/OrbitalJin/michi/internal/service"
	"github.com/OrbitalJin/michi/internal/store"
)

type Server struct {
	queryParser parser.QueryParserIface
	handler     handler.HandlerIface
	router      router.RouterIface
	services    *service.Services
	store       *store.Store
	config      *Config
}

func New(config *Config) (*Server, error) {
	store, err := store.New(config.storeCfg)

	if err != nil {
		return nil, fmt.Errorf("failed to access the store: %w", err)
	}

	if err = store.Migrate(); err != nil {
		return nil, fmt.Errorf("failed to conduct database migration: %w", err)
	}

	qp, err := parser.NewQueryParser(
		config.bangParserCfg,
		config.shortcutParserCfg,
		config.seshParserCfg,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create parser: %w", err)
	}

	psvc := service.NewSearchProviderService(
		qp.BangParser(),
		store.SearchProviders,
		config.serviceCgf,
	)
	hsvc := service.NewHistoryService(store.History)
	scsvc := service.NewShortcutService(store.Shortcuts)
	seshsvc := service.NewSessionService(store.Sessions)
	services := service.NewServices(psvc, hsvc, seshsvc, scsvc)

	handler := handler.NewHandler(qp, services, "q")

	router, err := router.NewRouter(handler)

	if err != nil {
		return nil, fmt.Errorf("failed to create router: %w", err)
	}

	router.Route()

	return &Server{
		queryParser: qp,
		store:       store,
		services:    services,
		router:      router,
		handler:     handler,
		config:      config,
	}, nil
}

func (server *Server) GetServices() *service.Services {
	return server.services
}

func (server *Server) Serve() {
	server.router.Serve(server.config.port)
}
