package productTransactions

import (
	"lapakUmkm/features/productTransactionDetails"
	"lapakUmkm/features/users"
)

type ProductTransactionEntity struct {
	Id                       uint
	UserId                   uint
	User                     users.UserEntity
	TotalProduct             int
	TotalPayment             int
	OrderId                  string
	PaymentStatus            string
	PaymentLink              string
	ProductTransactionDetail []productTransactionDetails.ProductTransactionDetailEntity
}

type ProductTransactionServiceInterface interface {
	Create(transctionEntity ProductTransactionEntity) (ProductTransactionEntity, error)
	MyTransactionHistory(myId, userId uint) ([]ProductTransactionEntity, error)
	GetById(id uint) (ProductTransactionEntity, error)
	CallBackMidtrans(id uint, status string) error
}

type ProductTransactionDataInterface interface {
	Store(transctionEntity ProductTransactionEntity) (uint, error)
	SelectById(id uint) (ProductTransactionEntity, error)
	Edit(transactionEntity ProductTransactionEntity, id uint) error
	SelectAll(userId uint) ([]ProductTransactionEntity, error)
}
