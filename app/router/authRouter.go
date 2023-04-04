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

	// g.GET("/sso-get-url", handler.GetSSOGoogleUrl)
	// g.GET("/sso-response-callback", handler.GetSSOGoogleUrl)

	g.POST("/sso-response-callback", handler.LoginSSOGoogle)

	g.GET("/forget-password", handler.ForgetPassword)
	g.GET("/user-exist", handler.IsUserExist)
	g.POST("/new-password", handler.NewPassword)

	g.POST("/change-password", handler.ChangePassword, middlewares.Authentication)
}
