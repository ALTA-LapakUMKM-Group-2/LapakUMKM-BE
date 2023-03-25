package router

import (
	"lapakUmkm/app/middlewares"
	_feedbacksData "lapakUmkm/features/feedbacks/data"
	_feedbacksHandler "lapakUmkm/features/feedbacks/delivery"
	_feedbacksService "lapakUmkm/features/feedbacks/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FeedbackRouter(db *gorm.DB, e *echo.Echo) {
	data := _feedbacksData.New(db)
	service := _feedbacksService.New(data)
	handler := _feedbacksHandler.New(service)

	e.GET("/feedbacks", handler.GetAll)
	e.GET("/feedbacks/:id", handler.GetById)
	e.GET("/products/:id/feedbacks", handler.GetFeedbackByProductId)

	g := e.Group("/feedbacks")
	g.POST("", handler.Create, middlewares.Authentication)
	g.PUT("/:id", handler.Update, middlewares.Authentication)
	g.DELETE("/:id", handler.Delete, middlewares.Authentication)
}