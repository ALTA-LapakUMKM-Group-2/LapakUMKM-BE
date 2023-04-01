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

	e.GET("/rooms/:id/chats", handler.GetByRoomId)

	// get last message from each senders. contains their sender_id 
	e.GET("chats/users", handler.GetSenderUser, middlewares.Authentication)

	// get all messages from all users who sent message to me (login user id)
	e.GET("chats", handler.AllMessageToMe, middlewares.Authentication)

	g := e.Group("/chats")
	g.POST("", handler.Create, middlewares.Authentication)
}