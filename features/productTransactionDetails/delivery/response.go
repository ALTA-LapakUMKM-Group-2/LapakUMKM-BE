package delivery

import (
	"lapakUmkm/features/productTransactionDetails"
	"lapakUmkm/features/products/delivery"
	"reflect"
)

type ProductTransactionDetailResponse struct {
	Id                   uint                     `json:"id"`
	ProductTransactionID uint                     `json:"product_transaction_id"`
	ProductId            uint                     `json:"product_id"`
	Product              delivery.ProductResponse `json:"product"`
	TotalProduct         int                      `json:"total_product"`
}

func EntityToResponse(e productTransactionDetails.ProductTransactionDetailEntity) ProductTransactionDetailResponse {
	result := ProductTransactionDetailResponse{
		Id:                   e.Id,
		ProductTransactionID: e.ProductTransaction.Id,
		ProductId:            e.ProductId,
		TotalProduct:         e.TotalProduct,
	}

	if !reflect.ValueOf(e.Product).IsZero() {
		result.Product = delivery.ProductResponse{
			ProductName: e.Product.ProductName,
		}
	}

	return result
}

func ListEntityToResponse(e []productTransactionDetails.ProductTransactionDetailEntity) []ProductTransactionDetailResponse {
	var dataResponses []ProductTransactionDetailResponse
	for _, v := range e {
		dataResponses = append(dataResponses, EntityToResponse(v))
	}
	return dataResponses
}
