package data

import (
	user "lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserId      uint
	User        *user.User `gorm:"foreignKey:UserId"`
	ProductName string
	Description string
	Price       int
	StockTotal  int
	StockSold   int
}
