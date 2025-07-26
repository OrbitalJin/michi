package models

import "time"

type SearchHistoryEvent struct {
	ID          int       `json:"id"`
	Query       string    `json:"query"`
	ProviderID  int       `json:"provider_id"`
	ProviderTag string    `json:"provider_tag"`
	Timestamp   time.Time `json:"timestamp"`
}
