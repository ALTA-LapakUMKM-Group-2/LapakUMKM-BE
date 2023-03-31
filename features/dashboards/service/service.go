package service

import (
	"encoding/json"
	"lapakUmkm/features/dashboards"
	"lapakUmkm/utils/helpers"
	"runtime"
	"strconv"

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
	redisName := "dashboardSS" + strconv.Itoa(int(userId))
	valueRedis, flag, err := helpers.GetRedis(redisName)
	if err != nil && !flag {
		return dashboards.DashboardEntity{}, err
	}

	//empty redis
	if err != nil && flag {
		runtime.GOMAXPROCS(2)
		if err := s.data.Create(userId); err != nil {
			return dashboards.DashboardEntity{}, err
		}
		data, _ := s.data.SelectByUserId(userId)

		var setRedis = func() {
			jsonByte, _ := json.Marshal(data)
			helpers.SetRedis(redisName, jsonByte)
		}

		go setRedis()

		return data, nil
	}

	var data dashboards.DashboardEntity
	err = json.Unmarshal([]byte(valueRedis), &data)
	if err != nil {
		panic(err)
	}

	return data, nil
}
