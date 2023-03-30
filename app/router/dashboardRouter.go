package router

import (
	"lapakUmkm/app/middlewares"
	_discussionsData "lapakUmkm/features/dashboards/data"
	_discussionsHandler "lapakUmkm/features/dashboards/delivery"
	_discussionsService "lapakUmkm/features/dashboards/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DashboardRouter(db *gorm.DB, e *echo.Echo) {
	data := _discussionsData.New(db)
	service := _discussionsService.New(data)
	handler := _discussionsHandler.New(service)

	e.GET("/dashboard", handler.Get, middlewares.Authentication)
	e.GET("/dashboard/create", handler.Create, middlewares.Authentication)
}
