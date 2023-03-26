package productTransactions

import "lapakUmkm/features/users"

type ProductTransactionEntity struct {
	Id     uint
	UserId uint
	User   users.UserEntity

}

type ProductTransactionServiceInterface interface {
	Create(transctionEntity ProductTransactionEntity) (ProductTransactionEntity, error)

}

type ProductTransactionDataInterface interface {
	Store(transctionEntity ProductTransactionEntity) (uint, error)
}