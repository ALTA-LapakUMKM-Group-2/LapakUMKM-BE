package delivery

import (
	"lapakUmkm/features/chats"
	user "lapakUmkm/features/users/delivery"
	"reflect"
)

type ChatResponse struct {
	Id          uint              `json:"id"`
	RoomId      string            `json:"room_id"`
	SenderId    uint              `json:"sender_id"`
	Sender      user.UserResponse `json:"sender"`
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
	if !reflect.ValueOf(chatEntity.Sender).IsZero() {
		chatResponse.Sender = user.UserResponse{
			FullName:     chatEntity.Sender.FullName,
			PhotoProfile: chatEntity.Sender.PhotoProfile,
		}
	}
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

type UserChatResponse struct {
	SenderId uint              `json:"sender_id"`
	Sender   user.UserResponse `json:"sender"`
	Text     string            `json:"text"`
}

func ToUserChatResponse(c chats.ChatEntity) UserChatResponse {
	userchat := UserChatResponse{
		SenderId: c.SenderId,
	}
	if !reflect.ValueOf(c.Sender).IsZero() {
		userchat.Sender = user.UserResponse{
			FullName:     c.Sender.FullName,
			PhotoProfile: c.Sender.PhotoProfile,
		}
	}
	return userchat
}

func ListUserResponse(c []chats.ChatEntity) []UserChatResponse {
	var dataRes []UserChatResponse
	for _, v := range c {
		dataRes = append(dataRes, ToUserChatResponse(v))
	}
	return dataRes
}

type TypeResponseChatWithMe struct {
	UserId       uint   `json:"user_id"`
	RoomId       string `json:"room_id"`
	FullName     string `json:"full_name"`
	PhotoProfile string `json:"photo_profile"`
}

func ListToResponseChat(chatEntity []chats.ChatEntity, userId uint) []TypeResponseChatWithMe {
	var dataRes []TypeResponseChatWithMe
	flag := map[string]bool{}

	for _, v := range chatEntity {
		_, ok := flag[v.RoomId]
		if ok {
			continue
		}
		flag[v.RoomId] = true

		if userId == v.SenderId {
			r := TypeResponseChatWithMe{
				UserId:       v.RecipientId,
				RoomId:       v.RoomId,
				FullName:     v.Recipient.FullName,
				PhotoProfile: v.Recipient.PhotoProfile,
			}

			dataRes = append(dataRes, r)
		} else {
			r := TypeResponseChatWithMe{
				UserId:       v.SenderId,
				RoomId:       v.RoomId,
				FullName:     v.Sender.FullName,
				PhotoProfile: v.Sender.PhotoProfile,
			}

			dataRes = append(dataRes, r)
		}
	}
	return dataRes
}
