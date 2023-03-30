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
	var dashboard Dashboard
	dashboard.UserId = userId
	if err := q.db.Create(&dashboard); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) SelectByUserId(id uint) (dashboards.DashboardEntity, error) {
	var dashboard Dashboard
	if err := q.db.First(&dashboard, id); err.Error != nil {
		return dashboards.DashboardEntity{}, err.Error
	}

	return ModelToEntity(dashboard), nil
}

func (q *query) Update(userId uint) error {
	var products data.Product
	q.db.Select("sum(product_transaction_details.total_product) as price,products.id, products.product_name").
		InnerJoins("ProductTransactionDetail").
		InnerJoins("ProductTransaction").
		Where("products.user_id = ?", userId).
		Group("product.id").
		Order("price desc").
		First(&products)

	var dashboard, update Dashboard
	update.FavoriteProductNameInWeek = products.ProductName
	update.TotalProductNameInWeek = uint(products.Price)

	q.db.Select("sum(product_transaction_details.total_product) as price").
		InnerJoins("ProductTransactionDetail").
		InnerJoins("ProductTransaction").
		Where("products.user_id = ?", userId).
		First(&products)
	update.TotalSellInWeek = uint(products.Price)

	q.db.Select("sum(product_transaction.total_payment) as price").
		InnerJoins("ProductTransactionDetail").
		InnerJoins("ProductTransaction").
		Where("products.user_id = ?", userId).
		First(&products)
	update.TotalCashInWeek = uint(products.Price)

	err := q.db.Model(&dashboard).
		Where("user_id = ?", userId).
		Updates(update)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
