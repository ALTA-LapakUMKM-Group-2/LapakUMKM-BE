package delivery

import "lapakUmkm/features/products"

type ProductRequest struct {
	CategoryId     uint   `json:"category_id" form:"category_id"`
	ProductName    string `json:"product_name" form:"product_name"`
	Description    string `json:"description" form:"description"`
	Price          int    `json:"price" form:"price"`
	StockRemaining int    `json:"stock_remaining" form:"stock_remaining"`
}

func ProductRequestToProductEntity(productRequest *ProductRequest) products.ProductEntity {
	return products.ProductEntity{
		CategoryId:     productRequest.CategoryId,
		ProductName:    productRequest.ProductName,
		Description:    productRequest.Description,
		Price:          productRequest.Price,
		StockRemaining: productRequest.StockRemaining,
	}
}
