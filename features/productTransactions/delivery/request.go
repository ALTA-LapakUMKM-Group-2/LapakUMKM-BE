package delivery

import (
	"lapakUmkm/features/productTransactionDetails/delivery"
	"lapakUmkm/features/productTransactions"
)

type TransactionRequest struct {
	TotalProduct              int                                         `json:"total_product" form:"total_product"`
	TotalPayment              int                                         `json:"total_payment" form:"total_payment"`
	ProductTransactionDetails []delivery.ProductTransactionDetailsRequest `json:"product_detail" form:"product_detail"`
}

// type Product

func TransactionRequestToTransactionEntity(transactionRequest *TransactionRequest) productTransactions.ProductTransactionEntity {
	result := productTransactions.ProductTransactionEntity{
		TotalProduct: transactionRequest.TotalProduct,
		TotalPayment: transactionRequest.TotalPayment,
	}

	for _, v := range transactionRequest.ProductTransactionDetails {
		result.ProductTransactionDetail = append(result.ProductTransactionDetail, delivery.RequestToEntity(&v))
	}

	return result
}
