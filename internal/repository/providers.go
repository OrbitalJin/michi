package repository // This package now called `repository`

import (
	"database/sql"
	"fmt"

	"github.com/OrbitalJin/qmuxr/internal/models"
	"github.com/mattn/go-sqlite3"
)

type ProviderRepo struct {
	db *sql.DB
}

func NewProviderRepo(db *sql.DB) *ProviderRepo {
	return &ProviderRepo{db: db}
}

func (r *ProviderRepo) Migrate() error {
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
	_, err := r.db.Exec(stmt)
	return err
}

func (r *ProviderRepo) GetByTag(tag string) (*models.SearchProvider, error) {
	var sp models.SearchProvider
	stmt := `
	SELECT id, tag, url, category, domain, rank, site_name, subcategory
	FROM search_providers WHERE tag = ?`
	err := r.db.QueryRow(stmt, tag).Scan(
		&sp.ID, &sp.Tag, &sp.URL, &sp.Category, &sp.Domain, &sp.Rank, &sp.SiteName, &sp.Subcategory,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query tag `%s`: %w", tag, err)
	}
	return &sp, nil
}

func (r *ProviderRepo) Insert(sp models.SearchProvider) error {
	stmt := `
	INSERT INTO search_providers
	(tag, url, category, domain, rank, site_name, subcategory)
	VALUES (?, ?, ?, ?, ?, ?, ?);`
	_, err := r.db.Exec(stmt, sp.Tag, sp.URL, sp.Category, sp.Domain, sp.Rank, sp.SiteName, sp.Subcategory)
	if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
		return fmt.Errorf("duplicate provider tag '%s': %w", sp.Tag, err)
	}
	return err
}
