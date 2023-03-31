package delivery

import "lapakUmkm/features/chats"

type ChatRequest struct {
	RecipientId uint   `json:"recipient_id" form:"recipient_id"`
	Text        string `json:"text" form:"text"`
}

func ReqToEntity(chatReq *ChatRequest) chats.ChatEntity{
	return chats.ChatEntity{
		RecipientId: chatReq.RecipientId,
		Text: chatReq.Text,
	}
}
