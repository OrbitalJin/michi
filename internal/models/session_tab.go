package models

type SessionTab struct {
	ID        int    `json:"id"`
	SessionID int    `json:"session_id"`
	URL       string `json:"url"`
}
