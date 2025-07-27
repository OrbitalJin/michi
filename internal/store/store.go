package store

import (
	"database/sql"
	"fmt"

	"github.com/OrbitalJin/qmuxr/internal/repository"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db              *sql.DB
	SearchProviders *repository.ProviderRepo
	History         *repository.HistoryRepo
	Shortcuts       *repository.ShortcutsRepository
	cfg             *Config
}

func New(cfg *Config) (*Store, error) {
	if cfg.path == "" {
		return nil, fmt.Errorf("path to database cannot be empty")
	}

	db, err := sql.Open("sqlite3", cfg.path)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:              db,
		cfg:             cfg,
		SearchProviders: repository.NewProviderRepo(db),
		History:         repository.NewHistoryRepo(db),
		Shortcuts:       repository.NewShortcutsRepository(db),
	}, nil
}

func (s *Store) Migrate() error {
	if err := s.SearchProviders.Migrate(); err != nil {
		return err
	}

	if err := s.History.Migrate(); err != nil {
		return err
	}

	if err := s.Shortcuts.Migrate(); err != nil {
		return err
	}

	return nil
}

func (s *Store) Shutdown() {
	_ = s.db.Close()
}

func (s *Store) GetCfg() *Config {
	return s.cfg
}
