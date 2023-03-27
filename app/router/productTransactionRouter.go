package router

import (
	"lapakUmkm/app/middlewares"
	_transactionsData "lapakUmkm/features/productTransactions/data"
	_transactionsHandler "lapakUmkm/features/productTransactions/delivery"
	_transactionsService "lapakUmkm/features/productTransactions/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRouter(db *gorm.DB, e *echo.Echo) {
	data := _transactionsData.New(db)
	service := _transactionsService.New(data)
	handler := _transactionsHandler.New(service)

	// e.GET("/feedbacks", handler.MyAllFeedback)
	// e.GET("/feedbacks/:id", handler.GetById)
	// e.GET("/products/:id/feedbacks", handler.GetFeedbackByProductId)

	g := e.Group("/transactions")
	g.POST("", handler.Create, middlewares.Authentication)
	// g.PUT("/:id", handler.Update, middlewares.Authentication)
	// g.DELETE("/:id", handler.Delete, middlewares.Authentication)
}