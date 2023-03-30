package delivery

import (
	"lapakUmkm/features/dashboards"
	"lapakUmkm/features/users/delivery"
)

type DashboardResponse struct {
	Id                        uint                  `json:"id"`
	UserId                    uint                  `json:"user_id"`
	User                      delivery.UserResponse `json:"user,omitempty"`
	FavoriteProductNameInWeek string                `json:"favorite_product_name_in_week"`
	TotalProductNameInWeek    uint                  `json:"total_product_name_in_week"`
	TotalSellInWeek           uint                  `json:"total_sell_in_week"`
	TotalCashInWeek           uint                  `json:"total_cash_in_week"`
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
