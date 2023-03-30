package productTransactionDetails

import (
	"lapakUmkm/features/products"
)

type ProductTransactionDetailEntity struct {
	Id                   uint
	ProductTransactionID uint
	ProductId            uint
	Product              products.ProductEntity
	TotalProduct         int
	Address              string
	Rating               float64
}

type ProductTransactionDetailServiceInterface interface {
	GetById(id uint) (ProductTransactionDetailEntity, error)
	GetByTransaksiId(transaksiId uint) ([]ProductTransactionDetailEntity, error)
	Create(productTransactionDetailEntity ProductTransactionDetailEntity) (ProductTransactionDetailEntity, error)
}

type ProductTransactionDetailDataInterface interface {
	SelectById(id uint) (ProductTransactionDetailEntity, error)
	SelectByTransaksiId(productId uint) ([]ProductTransactionDetailEntity, error)
	Store(productTransactionDetailEntity ProductTransactionDetailEntity) (uint, error)
}
