package dashboards

type DashboardEntity struct {
	Id                        uint   `json:"id"`
	UserId                    uint   `json:"user_id"`
	FavoriteProductNameInWeek string `json:"favorite_product_name_in_week"`
	TotalProductNameInWeek    uint   `json:"total_product_name_in_week"`
	TotalSellInWeek           uint   `json:"total_sell_in_week"`
	TotalCashInWeek           uint   `json:"total_cash_in_week"`
}

type DashboardDataInterface interface {
	SelectByUserId(id uint) (DashboardEntity, error)
	Create(userId uint) error
}

type DashboardServiceInterface interface {
	GetByUserId(id uint) (DashboardEntity, error)
}
