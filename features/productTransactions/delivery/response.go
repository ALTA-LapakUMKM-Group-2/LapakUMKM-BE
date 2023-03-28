package delivery

import (
	"lapakUmkm/features/productTransactions"
	user "lapakUmkm/features/users/delivery"
	"reflect"
)

type TransactionResponse struct {
	Id            uint              `json:"id"`
	UserId        uint              `json:"user_id"`
	TotalProduct  int               `json:"total_product"`
	TotalPayment  int               `json:"total_payment"`
	PaymentStatus string            `json:"payment_status"`
	PaymentLink   string            `json:"payment_link"`
	User          user.UserResponse `json:"user"`
}

func TransactionEntityToTransactionResponse(transactionEntity productTransactions.ProductTransactionEntity) TransactionResponse {
	transactionResponse := TransactionResponse{
		Id:            transactionEntity.Id,
		UserId:        transactionEntity.UserId,
		TotalProduct:  transactionEntity.TotalProduct,
		TotalPayment:  transactionEntity.TotalPayment,
		PaymentStatus: transactionEntity.PaymentStatus,
		PaymentLink:   transactionEntity.PaymentLink,
	}
	if !reflect.ValueOf(transactionEntity.User).IsZero() {
		transactionResponse.User = user.UserResponse{
			Id:       transactionEntity.User.Id,
			FullName: transactionEntity.User.FullName,
			Email:    transactionEntity.User.Email,
		}
	}
	return transactionResponse
}

func ListTransactionToTransactionResponse(transactionEntity []productTransactions.ProductTransactionEntity) []TransactionResponse {
	var dataRes []TransactionResponse
	for _, v := range transactionEntity {
		dataRes = append(dataRes, TransactionEntityToTransactionResponse(v))
	}
	return dataRes
}
