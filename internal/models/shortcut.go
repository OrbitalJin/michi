package models

type Shortcut struct {
	ID    int    `json:"id"`
	Alias string `json:"alias"`
	URL   string `json:"url"`
}
