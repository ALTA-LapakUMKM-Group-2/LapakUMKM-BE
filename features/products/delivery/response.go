package delivery

import (
	category "lapakUmkm/features/categories/delivery"
	productsimage "lapakUmkm/features/productImages/delivery"
	"lapakUmkm/features/products"
	user "lapakUmkm/features/users/delivery"
	"reflect"
)

type ProductResponse struct {
	Id             uint                                  `json:"id"`
	UserId         uint                                  `json:"user_id"`
	CategoryId     uint                                  `json:"category_id"`
	ProductName    string                                `json:"product_name"`
	Description    string                                `json:"description"`
	Size           string                                `json:"size"`
	Price          int                                   `json:"price"`
	StockRemaining int                                   `json:"stock_remaining"`
	StockSold      int                                   `json:"stock_sold"`
	User           user.UserResponse                     `json:"user"`
	Category       category.CategoryResponse             `json:"category"`
	ProductImage   []productsimage.ProductImagesResponse `json:"product_image"`
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
	}

	if !reflect.ValueOf(productEntity.User).IsZero() {
		productResponse.User = user.UserResponse{
			FullName: productEntity.User.FullName,
			Address:  productEntity.User.Address,
		}
	}

	if !reflect.ValueOf(productEntity.Category).IsZero() {
		productResponse.Category = category.CategoryResponse{
			Category: productEntity.Category.Category,
		}
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
