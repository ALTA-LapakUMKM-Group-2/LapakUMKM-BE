package data

import (
	"lapakUmkm/features/carts"

	"gorm.io/gorm"
)

type CartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) carts.CartData {
	return &CartQuery{
		db: db,
	}
}

// Add implements carts.CartData
func (baq *CartQuery) Add(newCart carts.Core) (carts.Core, error) {
	data := CoreToCart(newCart)
	tx := baq.db.Create(&data)
	if tx.Error != nil {
		return carts.Core{}, tx.Error
	}
	return CartToCore(data), nil
}
