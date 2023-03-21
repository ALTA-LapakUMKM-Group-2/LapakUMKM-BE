package data

import (
	dataProduct "lapakUmkm/features/products/data"
	dataUser "lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

type ProductTransaction struct {
	gorm.Model
	UserId        uint
	User          *dataUser.User `gorm:"foreignKey:UserId"`
	ProductId     uint
	Product       *dataProduct.Product `gorm:"foreignKey:ProductId"`
	OrderId       string
	PaymentStatus string
	PaymentLink   string
	TotalProduct  int
	TotalPayment  int
}
