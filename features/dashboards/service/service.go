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

func (s *DashboardService) Create(userId uint) error {
	return s.data.Create(userId)
}

func (s *DashboardService) GetByUserId(userId uint) (dashboards.DashboardEntity, error) {
	if err := s.data.Create(userId); err != nil {
		return dashboards.DashboardEntity{}, nil
	}
	return s.data.SelectByUserId(userId)
}

func (s *DashboardService) UpdateData(userId uint) error {
	if err := s.data.Update(userId); err != nil {
		return err
	}
	return nil
}
