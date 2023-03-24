package router

import (
	"lapakUmkm/app/middlewares"
	_productsData "lapakUmkm/features/products/data"
	_productsHandler "lapakUmkm/features/products/delivery"
	_productsService "lapakUmkm/features/products/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRouter(db *gorm.DB, e *echo.Echo) {
	data := _productsData.New(db)
	service := _productsService.New(data)
	handler := _productsHandler.New(service)

	g := e.Group("/products")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)

	g.Use(middlewares.Authentication)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
