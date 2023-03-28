package data

import (
	"lapakUmkm/features/productTransactionDetails"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) productTransactionDetails.ProductTransactionDetailDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectByProductId(productId uint) ([]productTransactionDetails.ProductTransactionDetailEntity, error) {
	var data []ProductTransactionDetail
	if err := q.db.Preload("Product").Find(&data); err.Error != nil {
		return nil, err.Error
	}

	return ListModelToEntity(data), nil
}

func (q *query) Store(productTransactionDetailEntity productTransactionDetails.ProductTransactionDetailEntity) (uint, error) {
	productTransactionDetail := EntityToModel(productTransactionDetailEntity)
	if err := q.db.Create(&productTransactionDetail); err.Error != nil {
		return 0, err.Error
	}
	return productTransactionDetail.Id, nil
}
