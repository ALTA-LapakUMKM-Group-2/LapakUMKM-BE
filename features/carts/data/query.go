package data

import (
	"errors"
	"lapakUmkm/features/carts"
	product "lapakUmkm/features/products/data"

	"github.com/jinzhu/copier"
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
func (cq *CartQuery) Add(newCart carts.Core) (carts.Core, error) {
	data := CoreToCart(newCart)
	var cart Cart
	var product product.Product
	// Cek apakah user mencoba membeli produk milik sendiri
	if err := cq.db.First(&product, data.ProductId).Error; err != nil {
		return carts.Core{}, errors.New("not found")
	}
	if product.UserId == data.UserId {
		return carts.Core{}, errors.New("you cannot add your own product to the cart")
	}
	// Cek apakah produk sudah ada di cart
	cq.db.Where("user_id = ? AND product_id = ?", data.UserId, data.ProductId).First(&cart)
	if cart.ID != 0 {
		cart.ProductPcs += data.ProductPcs
		if cart.ProductPcs > int64(product.StockRemaining) {
			return carts.Core{}, errors.New("quantity exceeds available stock")
		}
		cart.SubTotal = cart.ProductPcs * int64(product.Price)
		cq.db.Save(&cart)
		copier.Copy(&data, &cart)
	} else {
		if data.ProductPcs > int64(product.StockRemaining) {
			return carts.Core{}, errors.New("quantity exceeds available stock")
		}
		data.SubTotal = data.ProductPcs * int64(product.Price)
		if err := cq.db.Create(&data).Error; err != nil {
			return carts.Core{}, err
		}
	}
	return CartToCore(data), nil
}

// MyCart implements carts.CartData
func (cq *CartQuery) MyCart(userID uint) ([]carts.Core, error) {
	tmp := []Cart{}
	tx := cq.db.Where("carts.user_id = ?", userID).
		Joins("JOIN products ON carts.product_id = products.id").
		Where("products.deleted_at IS NULL").
		Joins("JOIN users ON products.user_id = users.id").
		Where("users.deleted_at IS NULL").
		Joins("JOIN product_images ON carts.product_id = product_images.product_id").
		Select("carts.id, carts.user_id, carts.product_id, carts.product_pcs, products.product_name AS product_name, products.price AS product_price, COALESCE(MIN(CONCAT('https://storage.googleapis.com/images_lapak_umkm/product/', product_images.image)), null) AS product_image, users.shop_name AS lapak_name, users.address AS lapak_address, CONCAT('https://storage.googleapis.com/images_lapak_umkm/profile/', users.photo_profile) AS photo_profile").
		Group("carts.id").
		Find(&tmp)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ListCartToCore(tmp), nil
}

// Update implements carts.CartData
func (cq *CartQuery) Update(updateCart carts.Core) (carts.Core, error) {
	data := CoreToCart(updateCart)
	var product product.Product
	if err := cq.db.Joins("JOIN carts ON carts.product_id = products.id").Where("carts.id = ?", data.ID).Find(&product).Error; err != nil {
		return carts.Core{}, err
	}
	if data.ProductPcs > int64(product.StockRemaining) {
		return carts.Core{}, errors.New("quantity exceeds available stock")
	}
	data.SubTotal = data.ProductPcs * int64(product.Price)
	tx := cq.db.Model(&Cart{}).Where("id = ? AND user_id = ?", data.ID, data.UserId).Updates(&data)
	if tx.RowsAffected < 1 {
		return carts.Core{}, errors.New("data not found")
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
		return errors.New("data not found")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// CartByID implements carts.CartData
func (cq *CartQuery) CartByID(userID uint, cart []uint) ([]carts.Core, error) {
	tmp := []Cart{}
	tx := cq.db.Where("carts.user_id = ? AND carts.id IN ?", userID, cart).Select("carts.id, carts.user_id, carts.product_id, carts.product_pcs, products.product_name AS product_name, products.price AS product_price, COALESCE(MIN(CONCAT('https://storage.googleapis.com/images_lapak_umkm/product/', product_images.image)), null) AS product_image, users.shop_name AS lapak_name, users.address AS lapak_address, CONCAT('https://storage.googleapis.com/images_lapak_umkm/profile/', users.photo_profile) AS photo_profile").Joins("JOIN products ON carts.product_id = products.id").Joins("JOIN users ON products.user_id = users.id").Joins("JOIN product_images ON carts.product_id = product_images.product_id").Group("carts.id").Find(&tmp)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ListCartToCore(tmp), nil
}

// BuyNow implements carts.CartData
func (cq *CartQuery) BuyNow(input carts.Core) (carts.Core, error) {
	data := CoreToCart(input)
	var cart Cart
	var product product.Product
	// Cek apakah user mencoba membeli produk milik sendiri
	cq.db.First(&product, data.ProductId)
	if product.UserId == data.UserId {
		return carts.Core{}, errors.New("you cannot add your own product to the cart")
	}
	if data.ProductPcs > int64(product.StockRemaining) {
		return carts.Core{}, errors.New("quantity exceeds available stock")
	}
	copier.Copy(&cart, &data)
	cart.SubTotal = data.ProductPcs * int64(product.Price)
	tx := cq.db.Where("products.id = ?", data.ProductId).Select("products.product_name AS product_name, products.price AS product_price, COALESCE(MIN(CONCAT('https://storage.googleapis.com/images_lapak_umkm/product/', product_images.image)), null) AS product_image, users.shop_name AS lapak_name, users.address AS lapak_address, CONCAT('https://storage.googleapis.com/images_lapak_umkm/profile/', users.photo_profile) AS photo_profile").Joins("JOIN products ON carts.product_id = products.id").Joins("JOIN users ON products.user_id = users.id").Joins("JOIN product_images ON products.id = product_images.product_id").Group("products.id").Find(&cart)
	if tx.Error != nil {
		return carts.Core{}, tx.Error
	}
	return CartToCore(cart), nil
}