package models

import "time"

type SearchHistoryEntry struct {
	ID          int       `json:"id"`
	Query       string    `json:"Query"`
	ProviderID  int       `json:"provider_id"`
	ProviderTag string    `json:"provider_tag"`
	Timestamp   time.Time `json:"timestamp"`
}
