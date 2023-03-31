package data

import (
	"lapakUmkm/features/dashboards"
	"lapakUmkm/features/products/data"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) dashboards.DashboardDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) Create(userId uint) error {
	var products data.Product
	q.db.Select("sum(product_transaction_details.total_product) as price,products.id, products.product_name").
		Joins("inner join product_transaction_details on product_transaction_details.product_id = products.id").
		Where("products.user_id = ?", userId).
		Group("products.id").
		Order("price desc").
		First(&products)

	var dashboard Dashboard
	dashboard.UserId = userId
	dashboard.FavoriteProductNameInWeek = products.ProductName
	dashboard.TotalProductNameInWeek = uint(products.Price)

	q.db.Select("sum(product_transaction_details.total_product) as price").
		Joins("inner join product_transaction_details on product_transaction_details.product_id = products.id").
		Where("products.user_id = ?", userId).
		Order("price desc").
		First(&products)
	dashboard.TotalSellInWeek = uint(products.Price)

	q.db.Select("sum(product_transaction_details.total_product * products.price) as price").
		Joins("inner join product_transaction_details on product_transaction_details.product_id = products.id").
		Where("products.user_id = ?", userId).
		Order("price desc").
		First(&products)
	dashboard.TotalCashInWeek = uint(products.Price)

	if err := q.db.Create(&dashboard); err.Error != nil {
		return err.Error
	}

	return nil
}

func (q *query) SelectByUserId(id uint) (dashboards.DashboardEntity, error) {
	var dashboard Dashboard
	if err := q.db.Where("user_id", id).First(&dashboard); err.Error != nil {
		return dashboards.DashboardEntity{}, err.Error
	}

	return ModelToEntity(dashboard), nil
}
