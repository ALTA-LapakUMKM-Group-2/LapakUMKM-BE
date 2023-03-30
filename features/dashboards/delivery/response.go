package delivery

import (
	"lapakUmkm/features/dashboards"
	"lapakUmkm/features/users/delivery"
)

type DashboardResponse struct {
	Id                    uint
	UserId                uint
	User                  delivery.UserResponse
	FavoriteProductInWeek uint
	TotalSellInWeek       uint
	TotalCashInWeek       uint
}

func EntityToResponse(entity dashboards.DashboardEntity) DashboardResponse {
	result := DashboardResponse{
		Id:                    entity.Id,
		UserId:                entity.UserId,
		FavoriteProductInWeek: entity.FavoriteProductInWeek,
		TotalSellInWeek:       entity.TotalSellInWeek,
		TotalCashInWeek:       entity.TotalCashInWeek,
	}
	return result
}

func ListEntityToListResponse(list []dashboards.DashboardEntity) (response []DashboardResponse) {
	for _, v := range list {
		response = append(response, EntityToResponse(v))
	}
	return
}
