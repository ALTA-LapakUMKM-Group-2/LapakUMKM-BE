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

	e.GET("/transactions", handler.MyTransactionHistory, middlewares.Authentication)
	e.GET("/transactions/:id", handler.GetById)

	g := e.Group("/transactions")
	g.POST("", handler.Create, middlewares.Authentication)

	e.POST("reservations/midtrans/callback", handler.CallBackMidtrans)
}