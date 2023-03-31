package router

import (
	"lapakUmkm/app/middlewares"
	_chatsData "lapakUmkm/features/chats/data"
	_chatsHandler "lapakUmkm/features/chats/delivery"
	_chatsService "lapakUmkm/features/chats/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ChatRouter(db *gorm.DB, e *echo.Echo) {
	data := _chatsData.New(db)
	service := _chatsService.New(data)
	handler := _chatsHandler.New(service)

	// e.GET("/chats", handler.MyAllchat)
	// e.GET("/chats/:id", handler.GetById)
	e.GET("/rooms/:id/chats", handler.GetByRoomId)

	g := e.Group("/chats")
	g.POST("", handler.Create, middlewares.Authentication)
	// g.PUT("/:id", handler.Update, middlewares.Authentication)
	// g.DELETE("/:id", handler.Delete, middlewares.Authentication)
}