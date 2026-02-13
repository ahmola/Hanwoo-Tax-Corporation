package models

type BaseEntity struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
