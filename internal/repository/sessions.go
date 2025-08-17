package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/OrbitalJin/michi/internal/models"
)

type SessionsRepoIface interface {
	Migrate() error
	Insert(session *models.Session) error
	GetFromAlias(alias string) (*models.Session, error)
	GetAll() ([]models.Session, error)
	Update(session *models.Session) error
	Delete(id int) error
}

type SessionsRepo struct {
	db *sql.DB
}

func NewSessionsRepo(db *sql.DB) *SessionsRepo {
	return &SessionsRepo{db}
}

func (repo *SessionsRepo) Migrate() error {
	stmt := `
		CREATE TABLE IF NOT EXISTS sessions (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			alias TEXT NOT NULL UNIQUE,
			urls TEXT,
			created_at DATETIME NOT NULL
		);
	`
	_, err := repo.db.Exec(stmt)
	if err != nil {
		return fmt.Errorf("failed to migrate sessions table: %w", err)
	}
	return nil
}

func (repo *SessionsRepo) Insert(session *models.Session) error {
	stmt := `
		INSERT INTO sessions
		(alias, urls, created_at)
		VALUES (?, ?, ?);
	`

	jsonData, err := json.Marshal(session.URLs)

	if err != nil {
		return err
	}

	createdAt := time.Now().Format(time.RFC3339)

	result, err := repo.db.Exec(stmt, session.Alias, jsonData, createdAt)

	if err != nil {
		return fmt.Errorf("failed to insert session '%s': %w", session.Alias, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert ID for session '%s': %w", session.Alias, err)
	}
	session.ID = int(id)
	return nil
}

func (repo *SessionsRepo) GetFromAlias(alias string) (*models.Session, error) {
	var session models.Session
	stmt := `SELECT id, alias, urls, created_at FROM sessions WHERE alias = ?`

	var rawURLs sql.NullString
	var createdAtStr string

	err := repo.db.QueryRow(stmt, alias).Scan(
		&session.ID,
		&session.Alias,
		&rawURLs,
		&createdAtStr,
	)

	session.CreatedAt, err = time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created_at for session `%s`: %w", session.Alias, err)
	}

	if rawURLs.Valid {
		var urls []string
		err = json.Unmarshal([]byte(rawURLs.String), &urls)

		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal urls for session `%s`: %w", session.Alias, err)
		}

		session.URLs = urls
	} else {
		session.URLs = []string{}
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to query session by alias '%s': %w", alias, err)
	}

	return &session, nil
}

func (repo *SessionsRepo) GetAll() ([]models.Session, error) {
	stmt := `SELECT id, alias, urls, created_at FROM sessions`

	rows, err := repo.db.Query(stmt)

	if err != nil {
		return nil, fmt.Errorf("failed to query all sessions: %w", err)
	}

	defer rows.Close()

	var sessions []models.Session
	for rows.Next() {
		var session models.Session
		var rawURLs sql.NullString
		var createdAtStr string

		err := rows.Scan(
			&session.ID,
			&session.Alias,
			&rawURLs,
			&createdAtStr,
		)

		session.CreatedAt, err = time.Parse(time.RFC3339, createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse created_at for session `%s`: %w", session.Alias, err)
		}

		if rawURLs.Valid {
			var urls []string
			err = json.Unmarshal([]byte(rawURLs.String), &urls)

			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal urls for session `%s`: %w", session.Alias, err)
			}

			session.URLs = urls

		} else {
			session.URLs = []string{}
		}

		if err != nil {
			return nil, fmt.Errorf("failed to scan session row: %w", err)
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (repo *SessionsRepo) Update(session *models.Session) error {
	stmt := `
		UPDATE sessions
		SET urls = ?
		WHERE id = ?;
	`

	jsonData, err := json.Marshal(session.URLs)

	if err != nil {
		return err
	}

	_, err = repo.db.Exec(stmt, jsonData, session.ID)
	if err != nil {
		return fmt.Errorf("failed to update session ID %d: %w", session.ID, err)
	}
	return nil
}

func (repo *SessionsRepo) Delete(id int) error {
	stmt := `DELETE FROM sessions WHERE id = ?;`
	res, err := repo.db.Exec(stmt, id)
	if err != nil {
		return fmt.Errorf("failed to delete session ID %d: %w", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for delete session ID %d: %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no session found with ID %d to delete", id)
	}
	return nil
}
