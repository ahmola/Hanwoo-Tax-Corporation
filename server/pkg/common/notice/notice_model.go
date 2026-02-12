package notice

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common"

type Notice struct {
	common.BaseEntity
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	IsImportant bool   `json:"is_important"`
}
