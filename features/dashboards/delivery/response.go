package delivery

import (
	"lapakUmkm/features/dashboards"
	"lapakUmkm/features/users/delivery"
)

type DashboardResponse struct {
	Id                        uint
	UserId                    uint
	User                      delivery.UserResponse
	FavoriteProductNameInWeek string
	TotalProductNameInWeek    uint
	TotalSellInWeek           uint
	TotalCashInWeek           uint
}

func EntityToResponse(entity dashboards.DashboardEntity) DashboardResponse {
	result := DashboardResponse{
		Id:                        entity.Id,
		UserId:                    entity.UserId,
		FavoriteProductNameInWeek: entity.FavoriteProductNameInWeek,
		TotalProductNameInWeek:    entity.TotalProductNameInWeek,
		TotalSellInWeek:           entity.TotalSellInWeek,
		TotalCashInWeek:           entity.TotalCashInWeek,
	}
	return result
}

func ListEntityToListResponse(list []dashboards.DashboardEntity) (response []DashboardResponse) {
	for _, v := range list {
		response = append(response, EntityToResponse(v))
	}
	return
}
