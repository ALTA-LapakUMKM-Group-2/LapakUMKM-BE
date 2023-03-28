package delivery

import (
	category "lapakUmkm/features/categories/delivery"
	productsimage "lapakUmkm/features/productImages/delivery"
	"lapakUmkm/features/products"
	user "lapakUmkm/features/users/delivery"
	"reflect"
)

type ProductResponse struct {
	Id             uint                                  `json:"id,omitempty"`
	UserId         uint                                  `json:"user_id,omitempty"`
	CategoryId     uint                                  `json:"category_id,omitempty"`
	ProductName    string                                `json:"product_name,omitempty"`
	Description    string                                `json:"description,omitempty"`
	Size           string                                `json:"size,omitempty"`
	Price          int                                   `json:"price,omitempty"`
	StockRemaining int                                   `json:"stock_remaining,omitempty"`
	StockSold      int                                   `json:"stock_sold,omitempty"`
	User           user.UserResponse                     `json:"user,omitempty"`
	Category       category.CategoryResponse             `json:"category,omitempty"`
	ProductImage   []productsimage.ProductImagesResponse `json:"product_image,omitempty"`
	Rating         float64                               `json:"rating,omitempty"`
}

func ProductEntityToProductResponse(productEntity products.ProductEntity) ProductResponse {
	productResponse := ProductResponse{
		Id:             productEntity.Id,
		UserId:         productEntity.UserId,
		CategoryId:     productEntity.CategoryId,
		ProductName:    productEntity.ProductName,
		Description:    productEntity.Description,
		Price:          productEntity.Price,
		StockRemaining: productEntity.StockRemaining,
		StockSold:      productEntity.StockSold,
		Size:           productEntity.Size,
		Rating:         productEntity.Rating,
	}

	if !reflect.ValueOf(productEntity.User).IsZero() {
		productResponse.User = user.UserResponse{
			FullName:     productEntity.User.FullName,
			Address:      productEntity.User.Address,
			ShopName:     productEntity.User.ShopName,
			PhotoProfile: productEntity.User.PhotoProfile,
		}
	}

	if !reflect.ValueOf(productEntity.Category).IsZero() {
		productResponse.Category = category.CategoryResponse{
			Category: productEntity.Category.Category,
		}
	}

	for _, v := range productEntity.ProductImage {
		var image = productsimage.ProductImagesResponse{
			Id:    v.Id,
			Image: "https://storage.googleapis.com/images_lapak_umkm/product/" + v.Image,
		}
		productResponse.ProductImage = append(productResponse.ProductImage, image)
	}

	return productResponse
}

func ListProductEntityToProductResponse(productEntity []products.ProductEntity) []ProductResponse {
	var dataResponses []ProductResponse
	for _, v := range productEntity {
		dataResponses = append(dataResponses, ProductEntityToProductResponse(v))
	}
	return dataResponses
}
