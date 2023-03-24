package router

import (
	"lapakUmkm/app/middlewares"
	_productsData "lapakUmkm/features/productImages/data"
	_productsHandler "lapakUmkm/features/productImages/delivery"
	_productsService "lapakUmkm/features/productImages/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductImagesRouter(db *gorm.DB, e *echo.Echo) {
	data := _productsData.New(db)
	service := _productsService.New(data)
	handler := _productsHandler.New(service)

	g := e.Group("/products")
	g.Use(middlewares.Authentication)
	g.GET("/:id/images", handler.GetByProductId)
	g.POST("/:id/upload-photo", handler.Create)
	g.DELETE("/:id/delete-photo/:photo_id", handler.Delete)
}
