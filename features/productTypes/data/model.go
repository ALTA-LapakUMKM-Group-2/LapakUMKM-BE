package data

import (
	"lapakUmkm/features/products/data"

	"gorm.io/gorm"
)

type ProductType struct {
	gorm.Model
	ProductId   uint
	Product     *data.Product `gorm:"foreignKey:ProductId"`
	ProductType string
}
