package productTransactionDetails

import (
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/features/products"
)

type ProductTransactionDetailEntity struct {
	Id                   uint
	ProductTransactionId uint
	ProductTransaction   productTransactions.ProductTransactionEntity
	ProductId            uint
	Product              products.ProductEntity
	TotalProduct         int
}

type ProductTransactionDetailServiceInterface interface {
	GetByProductId(productId uint) (ProductTransactionDetailEntity, error)
	Create(productTransactionDetailEntity ProductTransactionDetailEntity) (ProductTransactionDetailEntity, error)
}

type ProductTransactionDetailDataInterface interface {
	SelectByProductId(productId uint) ([]ProductTransactionDetailEntity, error)
	Store(productTransactionDetailEntity ProductTransactionDetailEntity) (uint, error)
}
