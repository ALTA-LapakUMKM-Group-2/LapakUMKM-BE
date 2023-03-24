package data

import (
	"gorm.io/gorm"
)

type ProductImages struct {
	gorm.Model
	ProductID uint
	Image     string
}
