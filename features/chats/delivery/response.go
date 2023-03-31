package delivery

import (
	"lapakUmkm/features/chats"
	user "lapakUmkm/features/users/delivery"
	"reflect"
)

type ChatResponse struct {
	Id       uint `json:"id"`
	RoomId   string `json:"room_id"`
	SenderId uint `json:"sender_id"`
	// Sender      user.UserResponse `json:"sender"`
	RecipientId uint              `json:"recipient_id"`
	Recipient   user.UserResponse `json:"recipient"`
	Text        string            `json:"text"`
}

func EntityToResponse(chatEntity chats.ChatEntity) ChatResponse {
	chatResponse := ChatResponse{
		Id:          chatEntity.Id,
		RoomId:      chatEntity.RoomId,
		SenderId:    chatEntity.SenderId,
		RecipientId: chatEntity.RecipientId,
		Text:        chatEntity.Text,
	}
	// if !reflect.ValueOf(chatEntity.Sender).IsZero(){
	// 	chatResponse.Sender = user.UserResponse{
	// 		FullName: chatEntity.Sender.FullName,
	// 		PhotoProfile: chatEntity.Sender.PhotoProfile,
	// 	}
	// }
	if !reflect.ValueOf(chatEntity.Recipient).IsZero() {
		chatResponse.Recipient = user.UserResponse{
			FullName:     chatEntity.Recipient.FullName,
			PhotoProfile: chatEntity.Recipient.PhotoProfile,
		}
	}
	return chatResponse
}

func ListEntityToResponse(chatEntity []chats.ChatEntity) []ChatResponse {
	var dataRes []ChatResponse
	for _, v := range chatEntity {
		dataRes = append(dataRes, EntityToResponse(v))
	}
	return dataRes
}
