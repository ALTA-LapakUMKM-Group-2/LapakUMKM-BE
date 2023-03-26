package productTransactions

import (
	"lapakUmkm/features/products"
	"lapakUmkm/features/users"
)

type ProductTransactionEntity struct {
	Id            uint
	UserId        uint
	User          users.UserEntity
	ProductId     uint
	Product       products.ProductEntity
	TotalProduct  int
	TotalPayment  int
	OrderId       string
	PaymentStatus string
	PaymentLink   string
}

type ProductTransactionServiceInterface interface {
	Create(transctionEntity ProductTransactionEntity) (ProductTransactionEntity, error)
}

type ProductTransactionDataInterface interface {
	Store(transctionEntity ProductTransactionEntity) (uint, error)
	SelectById(id uint) (ProductTransactionEntity, error)
	SelectProductPcs(transctionEntity ProductTransactionEntity) ([]ProductTransactionEntity, error)
	Edit(transactionEntity ProductTransactionEntity, id uint) error
}
