package service

import (
	"errors"
	"lapakUmkm/features/chats"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	Data chats.DataInterface
	validate *validator.Validate
}

func New(data chats.DataInterface) chats.ServiceInterface{
	return &Service{
		Data: data,
		validate: validator.New(),
	}
}

func (s *Service) Create(chatEntity chats.ChatEntity) (chats.ChatEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(chatEntity, "Sender", "Recipient")
	if errValidate != nil {
		return chats.ChatEntity{}, errValidate
	}
	chatId, err := s.Data.Store(chatEntity)
	if err != nil {
		return chats.ChatEntity{}, err
	}
	return s.Data.SelectById(chatId)
}

func (s *Service) GetByRoomId(roomId string) ([]chats.ChatEntity, error){
	res, err := s.Data.SelectByRoomId(roomId)
	if err != nil {
		// log.Println("query error", err.Error())
		return []chats.ChatEntity{}, errors.New("internal problem")
	}
	return res, nil
}

func (s *Service) GetById(id uint) (chats.ChatEntity, error) {
	return s.Data.SelectById(id)
}