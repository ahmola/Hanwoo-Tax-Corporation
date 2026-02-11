package common

type BaseEntity struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type Notice struct {
	BaseEntity
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	IsImportant bool   `json:"is_important"`
}

type Document struct {
	BaseEntity
	Title    string `json:"title"`
	Category string `json:"category"`
	FileUrl  string `json:"file_url`
	Size     string `json:"size"`
}

type Contact struct {
	BaseEntity
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
