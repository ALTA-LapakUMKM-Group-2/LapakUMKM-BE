package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"lapakUmkm/app/middlewares"
	_userData "lapakUmkm/features/users/data"
	_userHandler "lapakUmkm/features/users/delivery"
	_userService "lapakUmkm/features/users/service"
)

func UserRouter(db *gorm.DB, e *echo.Echo) {
	data := _userData.New(db)
	service := _userService.New(data)
	handler := _userHandler.New(service)

	g := e.Group("/users")
	g.Use(middlewares.Authentication)
	g.GET("", handler.GetUser)
	g.POST("", handler.Update)
	g.DELETE("", handler.Delete)
	g.POST("/update-to-seller", handler.UpdateToSeller)

	// g.POST("/photo-profile", handler.Update)
	//
}
