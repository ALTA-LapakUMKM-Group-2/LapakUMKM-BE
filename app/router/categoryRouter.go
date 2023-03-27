package router

import (
	"lapakUmkm/app/middlewares"
	categoryData "lapakUmkm/features/categories/data"
	categoryHandler "lapakUmkm/features/categories/delivery"
	categoryService "lapakUmkm/features/categories/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategoryRouter(db *gorm.DB, e *echo.Echo) {
	data := categoryData.New(db)
	service := categoryService.New(data)
	handler := categoryHandler.New(service)

	g := e.Group("/categories")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create, middlewares.Authentication)
	g.PUT("/:id", handler.Update, middlewares.Authentication)
	g.DELETE("/:id", handler.Delete, middlewares.Authentication)
}
