package productImages

import (
	"mime/multipart"
	"time"
)

type ProductImagesEntity struct {
	Id        uint
	ProductId uint
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductServiceInterface interface {
	GetByProductId(productId uint) ([]ProductImagesEntity, error)
	Create(productId uint, file *multipart.FileHeader) (ProductImagesEntity, error)
	Delete(productId uint, id uint) error
}

type ProductDataInterface interface {
	SelectById(productId uint, id uint) (ProductImagesEntity, error)
	SelectByProductId(productId uint) ([]ProductImagesEntity, error)
	Store(productImagesEntity ProductImagesEntity) (uint, error)
	Destroy(id uint) error
}
