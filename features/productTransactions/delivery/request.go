package delivery

import "lapakUmkm/features/productTransactions"

type TransactionRequest struct {
	// UserId       uint `json:"user_id" form:"user_id"`
	ProductId    uint `json:"product_id" form:"product_id"`
	TotalProduct int  `json:"total_product" form:"total_product"`
	TotalPayment int  `json:"total_payment" form:"total_payment"`
}

func TransactionRequestToTransactionEntity(transactionRequest *TransactionRequest) productTransactions.ProductTransactionEntity {
	return productTransactions.ProductTransactionEntity{
		// UserId: transactionRequest.UserId,
		ProductId: transactionRequest.ProductId,
		TotalProduct: transactionRequest.TotalProduct,
		TotalPayment: transactionRequest.TotalPayment,
	}
}