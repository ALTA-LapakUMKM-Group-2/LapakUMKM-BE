package service

import (
	"errors"
	"testing"

	"lapakUmkm/features/chats"
	"lapakUmkm/features/users"
	"lapakUmkm/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewDataInterface(t)
	inputData := chats.ChatEntity{SenderId: 1, RecipientId: 2, Text: "halo"}
	resData := chats.ChatEntity{Id: uint(1), SenderId: 1, RecipientId: 2, Sender: users.UserEntity{
		FullName: "jonathan", PhotoProfile: "",
	}, Recipient: users.UserEntity{
		FullName: "jajang", PhotoProfile: "",
	}, RoomId: "R12", Text: "HALO"}

	t.Run("success", func(t *testing.T) {
		repo.On("Store", mock.Anything).Return(uint(1), nil).Once()
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()

		srv := New(repo)
		res, err := srv.Create(inputData)
		assert.NoError(t, err)
		assert.Equal(t, resData.Id, res.Id)
		repo.AssertExpectations(t)
	})

	t.Run("errorValidation", func(t *testing.T) {
		repo.On("Store", mock.Anything).Return(uint(0), errors.New("required")).Once()
		srv := New(repo)
		_, err := srv.Create(inputData)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "required")
		repo.AssertExpectations(t)
	})

	t.Run("errorDuplicated", func(t *testing.T) {
		repo.On("Store", mock.Anything).Return(uint(0), errors.New("duplicated")).Once()
		srv := New(repo)
		_, err := srv.Create(resData)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "duplicated")
		repo.AssertExpectations(t)
	})

}

func TestGetByRoomId(t *testing.T) {
	repo := mocks.NewDataInterface(t)
	resData := []chats.ChatEntity{{Id: uint(1), SenderId: 1, RecipientId: 2, Sender: users.UserEntity{
		FullName: "jonathan", PhotoProfile: "",
	}, Recipient: users.UserEntity{
		FullName: "jajang", PhotoProfile: "",
	}, RoomId: "R12", Text: "HALO"}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectByRoomId", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		chats, err := srv.GetByRoomId("R12")
		assert.Nil(t, err)
		assert.Equal(t, resData[0].SenderId, chats[0].SenderId)
		repo.AssertExpectations(t)
	})
}

func TestGetSenderUser(t *testing.T) {
	repo := mocks.NewDataInterface(t)
	resData := []chats.ChatEntity{{Id: uint(1), SenderId: 1, RecipientId: 2, Sender: users.UserEntity{
		FullName: "jonathan", PhotoProfile: "",
	}, Recipient: users.UserEntity{
		FullName: "jajang", PhotoProfile: "",
	}, RoomId: "R12", Text: "HALO"}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		chats, err := srv.GetSenderUser(1, 3)
		assert.Nil(t, err)
		assert.Equal(t, resData[0].SenderId, chats[0].SenderId)
		repo.AssertExpectations(t)
	})
}

func TestAllMessageToMe(t *testing.T) {
	repo := mocks.NewDataInterface(t)
	resData := []chats.ChatEntity{{Id: uint(1), SenderId: 1, RecipientId: 2, Sender: users.UserEntity{
		FullName: "jonathan", PhotoProfile: "",
	}, Recipient: users.UserEntity{
		FullName: "jajang", PhotoProfile: "",
	}, RoomId: "R12", Text: "HALO"}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectAllMessageToMe", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		chats, err := srv.AllMessageToMe(1, 3)
		assert.Nil(t, err)
		assert.Equal(t, resData[0].SenderId, chats[0].SenderId)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := mocks.NewDataInterface(t)
	resData := chats.ChatEntity{Id: uint(1), SenderId: 1, RecipientId: 2, Sender: users.UserEntity{
		FullName: "jonathan", PhotoProfile: "",
	}, Recipient: users.UserEntity{
		FullName: "jajang", PhotoProfile: "",
	}, RoomId: "R12", Text: "HALO"}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		chatss, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, resData.SenderId, chatss.SenderId)
		repo.AssertExpectations(t)
	})
}
