package models

import "github.com/jinzhu/gorm"

// Good 商品模型
type Good struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Stock       uint16  `json:"stock"`
	Sales       uint16  `json:"sales"`
	Category    string  `json:"category"`
	PictureURL  string  `json:"pictureurl"`
	Description string  `json:"description"`
}
