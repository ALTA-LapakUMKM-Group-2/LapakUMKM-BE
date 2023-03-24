package data

import (
	"errors"
	"lapakUmkm/features/carts"
	"log"

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

// MyCart implements carts.CartData
func (cq *CartQuery) MyCart(userID uint) ([]carts.Core, error) {
	tmp := []Cart{}
	tx := cq.db.Where("carts.user_id = ?", userID).Select("carts.id, carts.user_id, carts.product_id, carts.product_pcs, products.product_name AS product_name, products.price AS product_price, users.full_name AS lapak_name, users.address AS lapak_address").Joins("JOIN users ON carts.user_id = users.id").Joins("JOIN products ON carts.product_id = products.id").Find(&tmp)

	if tx.Error != nil {
		return nil, tx.Error
	}
	log.Println(tmp)
	listFeedback := ListCartToCore(tmp)
	return listFeedback, nil
}

// Update implements carts.CartData
func (cq *CartQuery) Update(updateCart carts.Core) (carts.Core, error) {
	data := CoreToCart(updateCart)
	tx := cq.db.Model(&Cart{}).Where("id = ? AND user_id = ?", data.ID, data.UserId).Updates(&data)
	if tx.RowsAffected < 1 {
		return carts.Core{}, errors.New("no data updated")
	}
	if tx.Error != nil {
		return carts.Core{}, tx.Error
	}
	return CartToCore(data), nil
}

// Delete implements carts.CartData
func (cq *CartQuery) Delete(userID, cartID uint) error {
	tx := cq.db.Where("user_id = ?", userID).Delete(&Cart{}, cartID)
	if tx.RowsAffected < 1 {
		return errors.New("no data deleted")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
