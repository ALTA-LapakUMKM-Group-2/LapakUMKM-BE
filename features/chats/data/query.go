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
	if err := q.db.Where("room_id = ?", roomId).Preload("Recipient").Preload("Sender").Order("created_at asc").Find(&chat).Error; err != nil {
		return []chats.ChatEntity{}, err
	}
	return ListToEntity(chat), nil
}

func (q *query) SelectAll(userId uint) ([]chats.ChatEntity, error) {
    var chat []Chat
    if err := q.db.Joins(`
        INNER JOIN (
          SELECT sender_id, MAX(created_at) AS max_created_at
          FROM chats
          WHERE recipient_id = ?
          GROUP BY sender_id
        ) AS max_chats ON chats.sender_id = max_chats.sender_id AND chats.created_at = max_chats.max_created_at
    `, userId).Preload("Sender").Preload("Recipient").Order("created_at desc").Find(&chat).Error; err != nil {
        return []chats.ChatEntity{}, err
    }
    return ListToEntity(chat), nil
}

func (q *query) SelectAllMessageToMe(userId uint) ([]chats.ChatEntity, error) {
	chat := []Chat{}
	if err := q.db.Where("recipient_id = ?", userId).Preload("Sender").Preload("Recipient").Order("created_at desc").Find(&chat).Error; err != nil {
		return []chats.ChatEntity{}, err
	}
	return ListToEntity(chat), nil

}




// func (q *query) SelectAll(userId uint) ([]chats.ChatEntity, error) {
//     subquery := q.db.Table("chats").
//         Select("sender_id, max(id) as recent_chat_id").
//         Where("recipient_id = ?", userId).
//         Group("sender_id").SubQuery()

//     var chat []Chat
//     if err := q.db.Joins("JOIN (?) as sq on sq.sender_id = chats.sender_id AND sq.recent_chat_id = chats.id", subquery).
//         Preload("Sender").
//         Preload("Recipient").
//         Find(&chat).
//         Error; err != nil {
//         return []chats.ChatEntity{}, err
//     }
//     return ListToEntity(chat), nil
// }
