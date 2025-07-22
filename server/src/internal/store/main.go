package store

import (
	"database/sql"
	"fmt"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	conn *sql.DB
	cfg  *Config
}

func New(cfg *Config) (*Store, error) {
	if cfg.path == "" {
		return nil, fmt.Errorf("path to database cannot be empty")
	}

	conn, err := sql.Open("sqlite3", cfg.path)
	if err != nil {
		return nil, err
	}

	return &Store{
		conn: conn,
		cfg:  cfg,
	}, nil
}

func (s *Store) Migrate() error {
	stmt := `
	CREATE TABLE IF NOT EXISTS search_providers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tag TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL,
		category TEXT,
		domain TEXT,
		rank INTEGER,
		site_name TEXT,
		subcategory TEXT
	);`
	_, err := s.conn.Exec(stmt)
	return err
}

func (s *Store) GetProviderByTag(tag string) (*SearchProvider, error) {
	var sp SearchProvider

	stmt := `
	SELECT id, tag, url, category, domain, rank, site_name, subcategory
	FROM search_providers WHERE tag = ?`
	err := s.conn.QueryRow(stmt, tag).Scan(
		&sp.ID,
		&sp.Tag,
		&sp.URL,
		&sp.Category,
		&sp.Domain,
		&sp.Rank,
		&sp.SiteName,
		&sp.Subcategory,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to query tag `%s`: %w", tag, err)
	}

	return &sp, nil
}

func (s *Store) InsertProvider(sp SearchProvider) error {
	stmt := `
	INSERT INTO search_providers
	(tag, url, category, domain, rank, site_name, subcategory)
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err := s.conn.Exec(stmt,
		sp.Tag,
		sp.URL,
		sp.Category,
		sp.Domain,
		sp.Rank,
		sp.SiteName,
		sp.Subcategory,
	)

	if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
		fmt.Println(fmt.Errorf("skipping duplicate: %w", err))
	}

	return err
}

func (s *Store) Shutdown() {
	_ = s.conn.Close()
}
