package router

import (
	"lapakUmkm/app/middlewares"
	_transactionsData "lapakUmkm/features/productTransactions/data"
	_transactionsHandler "lapakUmkm/features/productTransactions/delivery"
	_transactionsService "lapakUmkm/features/productTransactions/service"

	_transactionDetailsData "lapakUmkm/features/productTransactionDetails/data"
	_transactionDetailsHandler "lapakUmkm/features/productTransactionDetails/delivery"
	_transactionDetailsService "lapakUmkm/features/productTransactionDetails/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRouter(db *gorm.DB, e *echo.Echo) {
	data := _transactionsData.New(db)
	service := _transactionsService.New(data)
	data2 := _transactionDetailsData.New(db)
	service2 := _transactionDetailsService.New(data2)
	handler := _transactionsHandler.New(service, service2)
	handler2 := _transactionDetailsHandler.New(service2)

	e.GET("/transactions", handler.MyTransactionHistory, middlewares.Authentication)
	e.GET("/transactions/:id", handler.GetById, middlewares.Authentication)
	e.GET("/transactions/:id/details", handler2.GetByProductId, middlewares.Authentication)

	g := e.Group("/transactions")
	g.POST("", handler.Create, middlewares.Authentication)

	e.POST("reservations/midtrans/callback", handler.CallBackMidtrans)
}
