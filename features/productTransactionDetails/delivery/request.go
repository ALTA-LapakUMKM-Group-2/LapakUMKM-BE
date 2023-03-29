package delivery

import (
	"lapakUmkm/features/productTransactionDetails"
)

type ProductTransactionDetailsRequest struct {
	ProductTransactionId uint   `json:"product_transaction_id" form:"product_transaction_id"`
	ProductId            uint   `json:"product_id" form:"product_id"`
	TotalProduct         int    `json:"total_product" form:"total_product"`
	Address              string `json:"address" form:"address"`
}

func RequestToEntity(r *ProductTransactionDetailsRequest) productTransactionDetails.ProductTransactionDetailEntity {
	return productTransactionDetails.ProductTransactionDetailEntity{
		ProductTransactionID: r.ProductTransactionId,
		ProductId:            r.ProductId,
		TotalProduct:         r.TotalProduct,
		Address:              r.Address,
	}
}
