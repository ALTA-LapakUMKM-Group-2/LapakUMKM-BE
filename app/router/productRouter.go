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

	e.GET("/products", handler.GetAll)
	e.GET("/products/:id", handler.GetById)

	g := e.Group("/products")
	g.POST("", handler.Create, middlewares.Authentication)
	g.PUT("/:id", handler.Update, middlewares.Authentication)
	g.DELETE("/:id", handler.Delete, middlewares.Authentication)
}
