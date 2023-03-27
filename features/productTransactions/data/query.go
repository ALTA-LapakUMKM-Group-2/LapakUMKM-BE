package data

import (
	"lapakUmkm/features/productTransactions"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) productTransactions.ProductTransactionDataInterface{
	return &query{
		db: db,
	}
}

func (qt *query) Store(transactionEntity productTransactions.ProductTransactionEntity) (uint, error) {
	transaction := TransactionEntityToTransaction(transactionEntity)
	if err := qt.db.Create(&transaction); err.Error != nil {
		return 0, err.Error
	}
	return transaction.ID, nil
}

func (qt *query) SelectById(id uint) (productTransactions.ProductTransactionEntity, error) {
	var transaction ProductTransaction
	if err := qt.db.Preload("User").Preload("Product").First(&transaction, id); err.Error != nil {
		return productTransactions.ProductTransactionEntity{}, err.Error
	}
	return TransactionToTransactionEntity(transaction), nil
}

func (qt *query) Edit(transactionEntity productTransactions.ProductTransactionEntity, id uint) error {
	transaction := TransactionEntityToTransaction(transactionEntity)
	if err := qt.db.Where("id", id).Updates(&transaction); err.Error != nil {
		return err.Error
	}
	return nil
}

func (qt *query) SelectProductPcs(transactionEntity productTransactions.ProductTransactionEntity) ([]productTransactions.ProductTransactionEntity, error) {
	var transactions []ProductTransaction
	err := qt.db.Where("product_id = ? AND total_product = ?", transactionEntity.ProductId, transactionEntity.TotalProduct).Find(&transactions)
	if err.Error != nil {
		return nil, err.Error
	}
	return ListTransactionToTransactionEntity(transactions), nil
}