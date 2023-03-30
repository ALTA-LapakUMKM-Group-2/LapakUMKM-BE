package dashboards

type DashboardEntity struct {
	Id                        uint
	UserId                    uint
	FavoriteProductNameInWeek string
	TotalProductNameInWeek    uint
	TotalSellInWeek           uint
	TotalCashInWeek           uint
}

type DashboardDataInterface interface {
	SelectByUserId(id uint) (DashboardEntity, error)
	Create(userId uint) error
	Update(userId uint) error
}

type DashboardServiceInterface interface {
	GetByUserId(id uint) (DashboardEntity, error)
	UpdateData(userId uint) error
}
