package data

import (
	"lapakUmkm/features/categories"
	category "lapakUmkm/features/categories/data"
	"lapakUmkm/features/productImages/data"
	"lapakUmkm/features/products"
	"lapakUmkm/features/users"
	user "lapakUmkm/features/users/data"
	"reflect"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserId         uint
	User           *user.User `gorm:"foreignKey:UserId"`
	CategoryId     uint
	Category       *category.Category `gorm:"foreignKey:CategoryId"`
	ProductName    string
	Description    string
	Size           string
	Price          int
	StockRemaining int
	StockSold      int
	ProductImage   []data.ProductImages
}

func ProductEntityToProduct(productEntity products.ProductEntity) Product {
	return Product{
		UserId:         productEntity.UserId,
		CategoryId:     productEntity.CategoryId,
		ProductName:    productEntity.ProductName,
		Description:    productEntity.Description,
		Price:          productEntity.Price,
		StockRemaining: productEntity.StockRemaining,
		StockSold:      productEntity.StockSold,
		Size:           productEntity.Size,
	}
}

func ProductToProductEntity(product Product) products.ProductEntity {
	result := products.ProductEntity{
		Id:             product.ID,
		UserId:         product.UserId,
		CategoryId:     product.CategoryId,
		ProductName:    product.ProductName,
		Description:    product.Description,
		Price:          product.Price,
		StockRemaining: product.StockRemaining,
		StockSold:      product.StockSold,
		Size:           product.Size,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}

	if !reflect.ValueOf(product.User).IsZero() {
		result.User = users.UserEntity{
			FullName: product.User.FullName,
			Address:  product.User.Address,
		}
	}

	if !reflect.ValueOf(product.Category).IsZero() {
		result.Category = categories.CategoryEntity{
			Category: product.Category.Category,
		}
	}

	return result
}

func ListProductToProductEntity(product []Product) []products.ProductEntity {
	var productEntity []products.ProductEntity
	for _, v := range product {
		productEntity = append(productEntity, ProductToProductEntity(v))
	}
	return productEntity
}
