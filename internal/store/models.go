package store

type SearchProvider struct {
	ID          int    `json:"id"`
	Category    string `json:"c"`
	Domain      string `json:"d"`
	Rank        int    `json:"r"`
	SiteName    string `json:"s"`
	Subcategory string `json:"sc"`
	Tag         string `json:"t"`
	URL         string `json:"u"`
}
