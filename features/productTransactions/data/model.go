package data

import (
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/features/products"
	dataProduct "lapakUmkm/features/products/data"
	"lapakUmkm/features/users"
	dataUser "lapakUmkm/features/users/data"
	"reflect"

	"gorm.io/gorm"
)

type ProductTransaction struct {
	gorm.Model
	UserId        uint
	User          *dataUser.User `gorm:"foreignKey:UserId"`
	ProductId     uint
	Product       *dataProduct.Product `gorm:"foreignKey:ProductId"`
	OrderId       string
	PaymentStatus string
	PaymentLink   string
	TotalProduct  int
	TotalPayment  int
}

func TransactionEntityToTransaction(transactionEntity productTransactions.ProductTransactionEntity) ProductTransaction {
	return ProductTransaction{
		UserId:        transactionEntity.Id,
		ProductId:     transactionEntity.ProductId,
		TotalProduct:  transactionEntity.TotalProduct,
		TotalPayment:  transactionEntity.TotalPayment,
		PaymentStatus: transactionEntity.PaymentStatus,
		PaymentLink:   transactionEntity.PaymentLink,
	}
}

func TransactionToTransactionEntity(transaction ProductTransaction) productTransactions.ProductTransactionEntity {
	result := productTransactions.ProductTransactionEntity{
		Id: transaction.ID,
		UserId: transaction.UserId,
		ProductId: transaction.ProductId,
		TotalProduct: transaction.TotalProduct,
		TotalPayment: transaction.TotalPayment,
		PaymentStatus: transaction.PaymentStatus,
		PaymentLink: transaction.PaymentLink,
	}
	if !reflect.ValueOf(transaction.User).IsZero() {
		result.User = users.UserEntity{
			FullName: transaction.User.FullName,
		}
	}
	if !reflect.ValueOf(transaction.Product).IsZero() {
		result.Product = products.ProductEntity{
			ProductName: transaction.Product.ProductName,
		}
	}
	return result
}
