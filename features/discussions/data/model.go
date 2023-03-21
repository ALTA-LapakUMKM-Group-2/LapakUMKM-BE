package data

import (
	"lapakUmkm/features/products/data"

	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	ProductId  uint
	Product    *data.Product `gorm:"foreignKey:ProductId"`
	ParentId   uint
	Discussion string
}
