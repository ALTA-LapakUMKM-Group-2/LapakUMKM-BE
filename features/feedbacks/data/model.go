package data

import (
	"lapakUmkm/features/products/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	ProductId uint
	Product   *data.Product `gorm:"foreignKey:ProductId"`
	ParentId  uint
	Feedback  string
}
