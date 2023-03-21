package data

import (
	dataProduct "lapakUmkm/features/products/data"
	dataUser "lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	UserId    uint
	User      *dataUser.User `gorm:"foreignKey:UserId"`
	ProductId uint
	Product   *dataProduct.Product `gorm:"foreignKey:ProductId"`
	Total     int
}
