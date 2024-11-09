package models

type URL struct {
	ID       string `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}
