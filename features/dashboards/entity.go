package dashboards

type DashboardEntity struct {
	Id                    uint
	UserId                uint
	FavoriteProductInWeek uint
	TotalSellInWeek       uint
	TotalCashInWeek       uint
}

type DashboardDataInterface interface {
	SelectByUserId(id uint) (DashboardEntity, error)
	UpdateFavoriteProductInWeek(userId, value uint) error
	UpdateTotalSellInWeek(userId, value uint) error
	UpdateTotalCashInWeek(userId, value uint) error
}

type DashboardServiceInterface interface {
	GetByUserId(id uint) (DashboardEntity, error)
	UpdateData(userId uint) error
}
