package router

import (
	"lapakUmkm/app/middlewares"
	_discussionsData "lapakUmkm/features/discussions/data"
	_discussionsHandler "lapakUmkm/features/discussions/delivery"
	_discussionsService "lapakUmkm/features/discussions/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DiscussionRouter(db *gorm.DB, e *echo.Echo) {
	data := _discussionsData.New(db)
	service := _discussionsService.New(data)
	handler := _discussionsHandler.New(service)

	e.GET("/discussions", handler.GetAll)
	e.GET("/discussions/:id", handler.GetById)
	e.GET("/products/:id/discussions", handler.GetDiscussionByProductId)

	g := e.Group("/discussions")
	g.POST("", handler.Create, middlewares.Authentication)
	g.PUT("/:id", handler.Update, middlewares.Authentication)
	g.DELETE("/:id", handler.Delete, middlewares.Authentication)
}
