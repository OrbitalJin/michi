package server

import (
	"fmt"

	"github.com/OrbitalJin/qmuxr/internal/parser"
	"github.com/OrbitalJin/qmuxr/internal/router"
	"github.com/OrbitalJin/qmuxr/internal/router/handler"
	"github.com/OrbitalJin/qmuxr/internal/service"
	"github.com/OrbitalJin/qmuxr/internal/store"
)

type Server struct {
	queryParser     parser.QueryParserIface
	providerService service.SPServiceIface
	historyService  service.HistoryServiceIface
	shortcutService service.ShortcutServiceIface
	handler         handler.HandlerIface
	router          router.RouterIface
	store           *store.Store
	config          *Config
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
	handler := handler.NewHandler(qp, psvc, hsvc, scsvc, seshsvc, "q")

	router, err := router.NewRouter(handler)

	if err != nil {
		return nil, fmt.Errorf("failed to create router: %w", err)
	}

	router.Route()

	return &Server{
		queryParser:     qp,
		providerService: psvc,
		historyService:  hsvc,
		shortcutService: scsvc,
		store:           store,
		router:          router,
		handler:         handler,
		config:          config,
	}, nil
}

func (server *Server) Start(port string) {
	server.router.Up(port)
}
