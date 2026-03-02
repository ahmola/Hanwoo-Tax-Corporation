package models

type Document struct {
	BaseEntity
	Title    string `json:"title"`
	Category string `json:"category"`
	FileUrl  string `json:"file_url"`
	Size     string `json:"size"`
}
