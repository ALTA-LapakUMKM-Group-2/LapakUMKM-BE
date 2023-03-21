package data

import (
	"lapakUmkm/features/categories/data"

	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	CategoryId  uint
	Category    *data.Category `gorm:"foreignKey:CategoryId"`
	SubCategory string
}
