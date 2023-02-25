package dataConfig

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
