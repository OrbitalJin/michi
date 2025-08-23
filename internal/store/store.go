package store

import (
	"database/sql"
	"fmt"

	"github.com/OrbitalJin/michi/internal/repository"
	_ "modernc.org/sqlite"
)

type Store struct {
	db              *sql.DB
	cfg             *Config
	SearchProviders repository.ProviderRepoIface
	History         repository.HistoryRepoIface
	Shortcuts       repository.ShortcutsRepoIface
	Sessions        repository.SessionsRepoIface
}

func New(cfg *Config) (*Store, error) {
	if cfg.path == "" {
		return nil, fmt.Errorf("path to database cannot be empty")
	}

	db, err := sql.Open("sqlite", cfg.path)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:              db,
		cfg:             cfg,
		SearchProviders: repository.NewProviderRepo(db),
		History:         repository.NewHistoryRepo(db),
		Shortcuts:       repository.NewShortcutsRepo(db),
		Sessions:        repository.NewSessionsRepo(db),
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

	if err := s.Sessions.Migrate(); err != nil {
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
