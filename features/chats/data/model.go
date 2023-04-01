package data

import (
	"fmt"
	"lapakUmkm/features/chats"
	"lapakUmkm/features/users"
	user "lapakUmkm/features/users/data"
	"reflect"

	// "strconv"
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	RoomId      string
	SenderId    uint
	RecipientId uint
	Sender      user.User `gorm:"foreignKey:SenderId"`
	Recipient   user.User `gorm:"foreignKey:RecipientId"`
	Text        string
	CreatedAt   time.Time
}

func (chat *Chat) BeforeCreate(tx *gorm.DB) error {
	var roomId string
	if chat.SenderId < chat.RecipientId {
		roomId = fmt.Sprintf("R%d%d", chat.SenderId, chat.RecipientId)
	} else {
		roomId = fmt.Sprintf("R%d%d", chat.RecipientId, chat.SenderId)
	}
	chat.RoomId = roomId
	return nil
}

func EntityToChat(chatEntity chats.ChatEntity) Chat {
	return Chat{
		RoomId:      chatEntity.RoomId,
		SenderId:    chatEntity.SenderId,
		RecipientId: chatEntity.RecipientId,
		Text:        chatEntity.Text,
		CreatedAt:   chatEntity.CreatedAt,
	}
}

func ChatToEntity(chat Chat) chats.ChatEntity {
	result := chats.ChatEntity{
		Id:          chat.ID,
		RoomId:      chat.RoomId,
		SenderId:    chat.SenderId,
		RecipientId: chat.RecipientId,
		Text:        chat.Text,
		CreatedAt:   chat.CreatedAt,
	}

	if !reflect.ValueOf(chat.Sender).IsZero() {
		result.Sender = users.UserEntity{
			FullName: chat.Sender.FullName,
			PhotoProfile: chat.Sender.PhotoProfile,
		}
	}

	if !reflect.ValueOf(chat.Recipient).IsZero() {
		result.Recipient = users.UserEntity{
			FullName:     chat.Recipient.FullName,
			PhotoProfile: chat.Recipient.PhotoProfile,
		}
	}
	return result
}

func ListToEntity(chat []Chat) []chats.ChatEntity {
	var chatEntity []chats.ChatEntity
	for _, v := range chat {
		chatEntity = append(chatEntity, ChatToEntity(v))
	}
	return chatEntity
}

