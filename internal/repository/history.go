package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/OrbitalJin/qmuxr/internal/models"
)

type HistoryRepoIface interface {
	Migrate() error
	Insert(entry *models.SearchHistoryEvent) error
	GetRecentHistory(limit int) ([]models.SearchHistoryEvent, error)
	DeleteEntry(id int) error
	DeleteOldHistory(beforeTime time.Time) error
}

type HistoryRepo struct {
	db *sql.DB
}

func NewHistoryRepo(db *sql.DB) *HistoryRepo {
	return &HistoryRepo{
		db: db,
	}
}

func (r *HistoryRepo) Migrate() error {
	stmt := `
	CREATE TABLE IF NOT EXISTS search_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		query TEXT NOT NULL,
		provider_id INTEGER,
		provider_tag TEXT,
		timestamp DATETIME NOT NULL
	);
	`
	_, err := r.db.Exec(stmt)
	return err
}

func (r *HistoryRepo) Insert(entry *models.SearchHistoryEvent) error {
	stmt := `
	INSERT INTO search_history
	(query, provider_id, provider_tag, timestamp)
	VALUES (?, ?, ?, ?);`

	_, err := r.db.Exec(stmt,
		entry.Query,
		entry.ProviderID,
		entry.ProviderTag,
		entry.Timestamp.Format(time.RFC3339),
	)

	return err
}

func (r *HistoryRepo) GetRecentHistory(limit int) ([]models.SearchHistoryEvent, error) {
	rows, err := r.db.Query(`
		SELECT id, query, provider_id, provider_tag, timestamp
		FROM search_history
		ORDER BY timestamp DESC
		LIMIT ?`,
		limit,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to query recent history: %w", err)
	}

	defer rows.Close()

	var history []models.SearchHistoryEvent
	for rows.Next() {
		var entry models.SearchHistoryEvent
		var timestampStr string

		err := rows.Scan(
			&entry.ID,
			&entry.Query,
			&entry.ProviderID,
			&entry.ProviderTag,
			&timestampStr,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan history row: %w", err)
		}

		entry.Timestamp, err = time.Parse(time.RFC3339, timestampStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse timestamp from DB: %w", err)
		}
		history = append(history, entry)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return history, nil
}

func (hr *HistoryRepo) DeleteEntry(id int) error {
	stmt := `DELETE FROM search_history WHERE id = ?;`
	_, err := hr.db.Exec(stmt, id)
	return err
}

func (hr *HistoryRepo) DeleteOldHistory(beforeTime time.Time) error {
	stmt := `DELETE FROM search_history WHERE timestamp < ?;`
	_, err := hr.db.Exec(stmt, beforeTime.Format(time.RFC3339))
	return err
}
