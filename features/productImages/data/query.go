package data

import (
	"lapakUmkm/features/productImages"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) productImages.ProductDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) Store(productImagesEntity productImages.ProductImagesEntity) (uint, error) {
	data := ProductImagesEntityToProductImages(productImagesEntity)
	if err := q.db.Create(&data); err.Error != nil {
		return 0, err.Error
	}
	return data.ID, nil
}

func (q *query) Destroy(id uint) error {
	var data ProductImages
	if err := q.db.Delete(&data, id); err.Error != nil {
		return err.Error
	}
	return nil
}

// SelectByProductId implements productsimages.ProductDataInterface
func (*query) SelectByProductId(productId uint) (productImages.ProductImagesEntity, error) {
	panic("unimplemented")
}
