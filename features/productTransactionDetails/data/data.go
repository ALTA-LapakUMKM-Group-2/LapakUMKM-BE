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

func (q *query) SelectById(id uint) (productTransactionDetails.ProductTransactionDetailEntity, error) {
	var data ProductTransactionDetail
	if err := q.db.Preload("Product").First(&data, id); err.Error != nil {
		return productTransactionDetails.ProductTransactionDetailEntity{}, err.Error
	}

	return ModelToEntity(data), nil
}

func (q *query) SelectByTransaksiId(productId uint) ([]productTransactionDetails.ProductTransactionDetailEntity, error) {
	var data []ProductTransactionDetail
	if err := q.db.Preload("Product").
		Select("product_transaction_details.id,product_transaction_details.product_transaction_id,product_transaction_details.product_id,product_transaction_details.total_product,CASE WHEN feedbacks.rating IS NULL THEN 0 ELSE feedbacks.rating END AS rating").
		Joins("left join feedbacks on feedbacks.product_transaction_detail_id = product_transaction_details.id").Where("product_transaction_id = ?", productId).Find(&data); err.Error != nil {
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
