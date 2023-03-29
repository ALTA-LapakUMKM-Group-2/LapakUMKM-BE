package data

import (
	"lapakUmkm/features/productTransactionDetails/data"
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/features/users"
	dataUser "lapakUmkm/features/users/data"
	"reflect"

	"gorm.io/gorm"
)

type ProductTransaction struct {
	gorm.Model
	UserId                   uint
	User                     *dataUser.User `gorm:"foreignKey:UserId"`
	OrderId                  string
	PaymentStatus            string
	PaymentLink              string
	TotalProduct             int
	TotalPayment             int
	ProductTransactionDetail []data.ProductTransactionDetail `gorm:"foreignKey:ProductTransactionID"`
}

func TransactionEntityToTransaction(transactionEntity productTransactions.ProductTransactionEntity) ProductTransaction {
	return ProductTransaction{
		UserId:        transactionEntity.UserId,
		TotalProduct:  transactionEntity.TotalProduct,
		TotalPayment:  transactionEntity.TotalPayment,
		PaymentStatus: transactionEntity.PaymentStatus,
		PaymentLink:   transactionEntity.PaymentLink,
		OrderId:       transactionEntity.OrderId,
	}
}

func TransactionToTransactionEntity(transaction ProductTransaction) productTransactions.ProductTransactionEntity {
	result := productTransactions.ProductTransactionEntity{
		Id:            transaction.ID,
		UserId:        transaction.UserId,
		TotalProduct:  transaction.TotalProduct,
		TotalPayment:  transaction.TotalPayment,
		PaymentStatus: transaction.PaymentStatus,
		PaymentLink:   transaction.PaymentLink,
		OrderId:       transaction.OrderId,
	}
	if !reflect.ValueOf(transaction.User).IsZero() {
		result.User = users.UserEntity{
			FullName: transaction.User.FullName,
		}
	}
	return result
}

func ListTransactionToTransactionEntity(transaction []ProductTransaction) []productTransactions.ProductTransactionEntity {
	var transactionEntity []productTransactions.ProductTransactionEntity
	for _, v := range transaction {
		transactionEntity = append(transactionEntity, TransactionToTransactionEntity(v))
	}
	return transactionEntity
}
