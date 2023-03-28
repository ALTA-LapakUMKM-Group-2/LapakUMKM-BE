package router

import (
	"lapakUmkm/app/middlewares"
	cartData "lapakUmkm/features/carts/data"
	cartHandler "lapakUmkm/features/carts/delivery"
	cartService "lapakUmkm/features/carts/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CartRouter(db *gorm.DB, e *echo.Echo) {
	data := cartData.New(db)
	service := cartService.New(data)
	handler := cartHandler.New(service)

	g := e.Group("/carts")
	g.Use(middlewares.Authentication)
	g.POST("", handler.Add)
	g.GET("", handler.MyCart)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
	e.GET("/checkout", handler.GetById)
	e.GET("/buynow", handler.BuyNow)
}
