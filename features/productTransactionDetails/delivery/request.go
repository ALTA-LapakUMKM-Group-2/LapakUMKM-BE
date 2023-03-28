package delivery

import "lapakUmkm/features/productTransactionDetails/data"

type ProductTransactionDetailsRequest struct {
	ProductTransactionId uint `json:"product_transaction_id" form:"product_transaction_id"`
	ProductId            uint `json:"product_id" form:"product_id"`
	TotalProduct         int  `json:"total_product" form:"total_product"`
}

func RequestToModel(r ProductTransactionDetailsRequest) data.ProductTransactionDetail {
	return data.ProductTransactionDetail{
		ProductTransactionID: r.ProductTransactionId,
		ProductId:            r.ProductId,
		TotalProduct:         r.TotalProduct,
	}
}
