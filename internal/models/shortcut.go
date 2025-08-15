package models

import "time"

type Shortcut struct {
	ID        int       `json:"id"`
	Alias     string    `json:"alias"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
