package database

type Bang struct {
	ID    int    `json:"id"`
	Alias string `json:"alias"`
	URL   string `json:"url"`
}
