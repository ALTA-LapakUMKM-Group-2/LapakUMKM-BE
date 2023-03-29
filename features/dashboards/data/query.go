package data

import (
	"lapakUmkm/features/dashboards"

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

func (q *query) SelectByUserId(id uint) (dashboards.DashboardEntity, error) {
	var dashboard Dashboard
	if err := q.db.First(&dashboard); err.Error != nil {
		return dashboards.DashboardEntity{}, err.Error
	}

	return ModelToEntity(dashboard), nil
}

func (q *query) UpdateFavoriteProductInWeek(userId, value uint) error {
	var dashboard Dashboard
	err := q.db.Model(&dashboard).
		Where("user_id = ?", userId).
		Update("favorite_product_in_week", value)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) UpdateTotalCashInWeek(userId uint, value uint) error {
	var dashboard Dashboard
	err := q.db.Model(&dashboard).
		Where("user_id = ?", userId).
		Update("total_cash_in_week", value)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) UpdateTotalSellInWeek(userId uint, value uint) error {
	var dashboard Dashboard
	err := q.db.Model(&dashboard).
		Where("user_id = ?", userId).
		Update("total_sell_in_week", value)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
