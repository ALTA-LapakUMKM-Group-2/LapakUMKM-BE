package chats

import (
	"lapakUmkm/features/users"
	"time"
)

// import firebase "firebase.google.com/go"

type ChatEntity struct {
	Id          uint
	RoomId      string
	SenderId    uint
	RecipientId uint `validate:"required"`
	Sender      users.UserEntity
	Recipient   users.UserEntity
	Text        string
	CreatedAt   time.Time
}

type ServiceInterface interface {
	Create(chatEntity ChatEntity) (ChatEntity, error)
	GetByRoomId(roomId string) ([]ChatEntity, error)
	GetById(id uint) (ChatEntity, error)
	GetSenderUser(myId, userId uint) ([]ChatEntity, error)
	AllMessageToMe(myId, userId uint) ([]ChatEntity, error)
}

type DataInterface interface {
	Store(chatEntity ChatEntity) (uint, error)
	SelectById(id uint) (ChatEntity, error)
	SelectByRoomId(roomId string) ([]ChatEntity, error)
	SelectAll(userId uint) ([]ChatEntity, error)
	SelectAllMessageToMe(userId uint) ([]ChatEntity, error)
}
