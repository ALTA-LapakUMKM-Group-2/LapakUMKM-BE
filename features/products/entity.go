package products

import (
	"lapakUmkm/features/categories"
	"lapakUmkm/features/productImages"
	"lapakUmkm/features/users"
	"time"
)

type ProductEntity struct {
	Id             uint
	UserId         uint
	CategoryId     uint
	ProductName    string
	Description    string
	Size           string
	Price          int
	StockRemaining int
	StockSold      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           users.UserEntity
	Category       categories.CategoryEntity
	ProductImage   []productImages.ProductImagesEntity
}

type ProductFilter struct {
	PriceMin   int
	PriceMax   int
	CategoryId uint
	Rating     float64
}

type ProductServiceInterface interface {
	GetAll(productFilter ProductFilter) ([]ProductEntity, error)
	GetById(id uint) (ProductEntity, error)
	Create(productEntity ProductEntity) (ProductEntity, error)
	Update(productEntity ProductEntity, id, userId uint) (ProductEntity, error)
	Delete(id, userId uint) error
	GetProductByUserId(userId uint)
	GetProductByCategoryId(userId uint)
}

type ProductDataInterface interface {
	SelectAll(productFilter ProductFilter) ([]ProductEntity, error)
	SelectById(id uint) (ProductEntity, error)
	Store(productEntity ProductEntity) (uint, error)
	Edit(productEntity ProductEntity, id uint) error
	Destroy(id uint) error
}
