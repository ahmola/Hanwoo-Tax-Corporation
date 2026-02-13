package models

type Contact struct {
	BaseEntity
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
