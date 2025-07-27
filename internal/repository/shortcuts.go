package repository

import (
	"database/sql"
	"fmt"

	"github.com/OrbitalJin/qmuxr/internal/models"
)

type ShortcutsRepoIface interface {
	Migrate() error
	Insert(shortcut *models.Shortcut) error
	GetFromAlias(alias string) (*models.Shortcut, error)
	Delete(id int) error
}

type ShortcutsRepo struct {
	db *sql.DB
}

func NewShortcutsRepo(db *sql.DB) *ShortcutsRepo {
	return &ShortcutsRepo{db}
}

func (repo *ShortcutsRepo) Migrate() error {
	stmt := `
		CREATE TABLE IF NOT EXISTS shortcuts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
		);
	`
	_, err := repo.db.Exec(stmt)
	return err
}

func (repo *ShortcutsRepo) Insert(shortcut *models.Shortcut) error {
	stmt := `
		INSERT INTO shortcuts 
		(alias, url)
		VALUES (?, ?);
	`
	_, err := repo.db.Exec(stmt, shortcut.Alias, shortcut.URL)
	return err
}

func (repo *ShortcutsRepo) GetFromAlias(alias string) (*models.Shortcut, error) {
	var shortcut models.Shortcut
	stmt := `SELECT id, alias, url FROM shortcuts WHERE alias = ?`

	err := repo.db.QueryRow(stmt, alias).Scan(
		&shortcut.ID,
		&shortcut.Alias,
		&shortcut.URL,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query alias `%s`: %w", alias, err)
	}

	return &shortcut, nil
}

func (repo *ShortcutsRepo) Delete(id int) error {
	stmt := `DELETE FROM shortcuts WHERE id = ?`

	_, err := repo.db.Exec(stmt, id)
	return err
}
