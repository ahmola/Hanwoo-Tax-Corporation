package contact

import "github.com/ahmola/Hanwoo-Tax-Corporation/pkg/common"

type Contact struct {
	common.BaseEntity
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
