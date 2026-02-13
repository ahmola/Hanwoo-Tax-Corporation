package models

type Notice struct {
	BaseEntity
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	IsImportant bool   `json:"is_important"`
}
