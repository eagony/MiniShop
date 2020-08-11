package models

import (
	"github.com/jinzhu/gorm"
)

// Order 订单模型
type Order struct {
	gorm.Model
	Detail  string `json:"detail"`
	Status  string `json:"status"`
	Amount  uint   `json:"amount"`
	Remark  string `json:"remark"`
	Address string `json:"address"`
}
