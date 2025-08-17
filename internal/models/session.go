package models

import "time"

type Session struct {
	ID        int       `json:"id"`
	Alias     string    `json:"alias"`
	URLs      []string  `json:"urls"`
	CreatedAt time.Time `json:"created_at"`
}
