package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/OrbitalJin/michi/internal/models"
)

type ShortcutsRepoIface interface {
	Migrate() error
	Insert(shortcut *models.Shortcut) error
	GetFromAlias(alias string) (*models.Shortcut, error)
	GetAll() ([]models.Shortcut, error)
	Delete(id int) error
	DeleteFromAlias(alias string) error
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
			url TEXT NOT NULL,
			created_at DATETIME NOT NULL
		);
	`
	_, err := repo.db.Exec(stmt)
	return err
}

func (repo *ShortcutsRepo) Insert(shortcut *models.Shortcut) error {
	stmt := `
		INSERT INTO shortcuts
		(alias, url, created_at)
		VALUES (?, ?, ?);
	`

	_, err := repo.db.Exec(
		stmt,
		shortcut.Alias,
		shortcut.URL,
		time.Now().Format(time.RFC3339),
	)
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

func (repo *ShortcutsRepo) GetAll() ([]models.Shortcut, error) {
	stmt := `SELECT id, alias, url, created_at FROM shortcuts`

	rows, err := repo.db.Query(stmt)

	if err != nil {
		return nil, fmt.Errorf("failed to query all shortcuts: %w", err)
	}

	defer rows.Close()

	var shortcuts []models.Shortcut
	for rows.Next() {
		var shortcut models.Shortcut
		var createdAtStr string

		err := rows.Scan(
			&shortcut.ID,
			&shortcut.Alias,
			&shortcut.URL,
			&createdAtStr,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan shortcut row: %w", err)
		}

		shortcut.CreatedAt, err = time.Parse(time.RFC3339, createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse created_at for shortcut `%s`: %w", shortcut.Alias, err)
		}

		shortcuts = append(shortcuts, shortcut)
	}

	return shortcuts, nil
}

func (repo *ShortcutsRepo) Delete(id int) error {
	stmt := `DELETE FROM shortcuts WHERE id = ?`

	_, err := repo.db.Exec(stmt, id)
	return err
}

func (repo *ShortcutsRepo) DeleteFromAlias(alias string) error {
	stmt := `DELETE FROM shortcuts WHERE alias = ?`

	_, err := repo.db.Exec(stmt, alias)
	return err
}
