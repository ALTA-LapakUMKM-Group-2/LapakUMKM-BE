package data

import (
	"lapakUmkm/features/chats"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) chats.DataInterface{
	return &query{
		db: db,
	}
}

func (q *query) Store(chatEntity chats.ChatEntity) (uint, error) {
	discussion := EntityToChat(chatEntity)
	if err := q.db.Create(&discussion); err.Error != nil {
		return 0, err.Error
	}
	return discussion.ID, nil
}

func (q *query) SelectById(id uint) (chats.ChatEntity, error) {
	var chat Chat
	if err := q.db.Preload("Sender").Preload("Recipient").First(&chat, id); err.Error != nil {
		return chats.ChatEntity{}, err.Error
	}

	return ChatToEntity(chat), nil
}

func (q *query) SelectByRoomId(roomId string) ([]chats.ChatEntity, error) {
	chat := []Chat{}
	if err := q.db.Where("room_id = ?", roomId).Preload("Sender").Preload("Recipient").Order("created_at asc").Find(&chat).Error; err != nil {
		return []chats.ChatEntity{}, err
	}
	return ListToEntity(chat), nil
}
