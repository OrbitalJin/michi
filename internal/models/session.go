package models

type Session struct {
	ID    int    `json:"id"`
	Alias string `json:"alias"`
	URLs  []SessionTab
}
