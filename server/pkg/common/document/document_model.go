package document

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common"

type Document struct {
	common.BaseEntity
	Title    string `json:"title"`
	Category string `json:"category"`
	FileUrl  string `json:"file_url`
	Size     string `json:"size"`
}
