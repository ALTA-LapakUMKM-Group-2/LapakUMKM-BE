package service

import (
	"lapakUmkm/features/dashboards"

	"github.com/go-playground/validator/v10"
)

type DashboardService struct {
	data     dashboards.DashboardDataInterface
	validate *validator.Validate
}

func New(data dashboards.DashboardDataInterface) dashboards.DashboardServiceInterface {
	return &DashboardService{
		data:     data,
		validate: validator.New(),
	}
}

func (s *DashboardService) GetByUserId(userId uint) (dashboards.DashboardEntity, error) {
	return s.data.SelectByUserId(userId)
}

func (s *DashboardService) UpdateData(userId uint) error {
	if err := s.data.UpdateFavoriteProductInWeek(userId, 10); err != nil {
		return err
	}

	return nil
}
