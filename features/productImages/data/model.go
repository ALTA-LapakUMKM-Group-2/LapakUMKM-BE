package data

import (
	"lapakUmkm/features/productImages"

	"gorm.io/gorm"
)

type ProductImages struct {
	gorm.Model
	ProductID uint
	Image     string
}

func ProductImagesEntityToProductImages(p productImages.ProductImagesEntity) ProductImages {
	return ProductImages{
		ProductID: p.ProductId,
		Image:     p.Image,
	}
}

func ProductImagesToProductImagesEntity(p ProductImages) productImages.ProductImagesEntity {
	result := productImages.ProductImagesEntity{
		Id:        p.ID,
		ProductId: p.ProductID,
		Image:     p.Image,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	return result
}

func ToListEntity(pImage []ProductImages) []productImages.ProductImagesEntity {
	var images []productImages.ProductImagesEntity
	for _, v := range pImage {
		images = append(images, ProductImagesToProductImagesEntity(v))
	}

	return images
}
