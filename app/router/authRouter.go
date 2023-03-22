package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"lapakUmkm/app/middlewares"
	_authData "lapakUmkm/features/auth/data"
	_authHandler "lapakUmkm/features/auth/delivery"
	_authService "lapakUmkm/features/auth/service"
)

func AuthRouter(db *gorm.DB, e *echo.Echo) {
	data := _authData.New(db)
	service := _authService.New(data)
	handler := _authHandler.New(service)

	g := e.Group("/auth")
	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
	// g.POST("/forget-password", handler.Create)

	g.Use(middlewares.Authentication)
	g.POST("/change-password", handler.ChangePassword)
}
