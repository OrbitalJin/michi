package repository

import (
	"database/sql"
	"fmt"
	"github.com/OrbitalJin/michi/internal/models"
	_ "modernc.org/sqlite"
)

type ProviderRepoIface interface {
	Migrate() error
	GetByTag(tag string) (*models.SearchProvider, error)
	GetAll() ([]models.SearchProvider, error)
	Insert(sp models.SearchProvider) error
	Delete(id int) error
}

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
	if err != nil {
		return fmt.Errorf("failed to insert provider %q: %w", sp.Tag, err)
	}
	return nil
}

func (r *ProviderRepo) GetAll() ([]models.SearchProvider, error) {
	stmt := `SELECT id, tag, url, category, domain, rank, site_name, subcategory FROM search_providers`
	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, fmt.Errorf("failed to query all providers: %w", err)
	}
	defer rows.Close()
	var providers []models.SearchProvider
	for rows.Next() {
		var sp models.SearchProvider
		err := rows.Scan(&sp.ID, &sp.Tag, &sp.URL, &sp.Category, &sp.Domain, &sp.Rank, &sp.SiteName, &sp.Subcategory)
		if err != nil {
			return nil, fmt.Errorf("failed to scan provider: %w", err)
		}
		providers = append(providers, sp)
	}
	return providers, nil
}

func (r *ProviderRepo) Delete(id int) error {
	stmt := `DELETE FROM search_providers WHERE id = ?`
	_, err := r.db.Exec(stmt, id)
	return err
}
