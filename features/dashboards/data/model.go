package data

import (
	"lapakUmkm/features/dashboards"
	"lapakUmkm/features/users/data"

	"gorm.io/gorm"
)

type Dashboard struct {
	gorm.Model
	UserId                uint
	User                  *data.User `gorm:"foreignKey:UserId"`
	FavoriteProductInWeek uint
	TotalSellInWeek       uint
	TotalCashInWeek       uint
}

func EntityToModel(entity dashboards.DashboardEntity) Dashboard {
	return Dashboard{
		UserId:                entity.UserId,
		FavoriteProductInWeek: entity.FavoriteProductInWeek,
		TotalSellInWeek:       entity.TotalSellInWeek,
		TotalCashInWeek:       entity.TotalCashInWeek,
	}
}

func ModelToEntity(model Dashboard) dashboards.DashboardEntity {
	return dashboards.DashboardEntity{
		Id:                    model.ID,
		UserId:                model.UserId,
		FavoriteProductInWeek: model.FavoriteProductInWeek,
		TotalSellInWeek:       model.TotalSellInWeek,
		TotalCashInWeek:       model.TotalCashInWeek,
	}
}

func ListModelToListEntity(model []Dashboard) (data []dashboards.DashboardEntity) {
	for _, v := range model {
		data = append(data, ModelToEntity(v))
	}
	return
}
