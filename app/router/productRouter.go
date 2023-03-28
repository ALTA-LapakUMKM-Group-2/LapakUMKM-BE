package router

import (
	"lapakUmkm/app/middlewares"
	_productsData "lapakUmkm/features/products/data"
	_productsHandler "lapakUmkm/features/products/delivery"
	_productsService "lapakUmkm/features/products/service"

	_productImagesData "lapakUmkm/features/productImages/data"
	_productImagesHandler "lapakUmkm/features/productImages/delivery"
	_productImagesService "lapakUmkm/features/productImages/service"

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

	data2 := _productImagesData.New(db)
	service2 := _productImagesService.New(data2)
	handler2 := _productImagesHandler.New(service2, service)

	g.GET("/:id/images", handler2.GetByProductId)
	g.POST("/:id/upload-photo", handler2.Create, middlewares.Authentication)
	g.DELETE("/:id/delete-photo/:photo_id", handler2.Delete, middlewares.Authentication)

}
