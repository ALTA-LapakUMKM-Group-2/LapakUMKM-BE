package service

import (
	"errors"
	"lapakUmkm/features/users"
	"lapakUmkm/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockData = users.UserEntity{
	Id:           uint(1),
	FullName:     "aa",
	Email:        "mail@mail.com",
	Password:     "123",
	PhoneNumber:  "",
	Address:      "",
	Role:         "user",
	ShopName:     "",
	PhotoProfile: "",
}

func TestGetUser(t *testing.T) {
	data := mocks.NewUserDataInterface(t)
	srv := New(data)
	t.Run("Success", func(t *testing.T) {
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		resultData, err := srv.GetUser(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, resultData.Id, mockData.Id)
		data.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	data := mocks.NewUserDataInterface(t)
	srv := New(data)
	var input = users.UserEntity{
		FullName: "aa",
		Email:    "mail@mail.com",
		Password: "123",
		Role:     "user",
	}
	t.Run("Success", func(t *testing.T) {
		data.On("Store", input).Return(uint(1), nil).Once()
		data.On("SelectById", uint(1)).Return(mockData, nil)
		resultData, err := srv.Create(input)
		assert.NoError(t, err)
		assert.Equal(t, resultData.Email, input.Email)
		data.AssertExpectations(t)
	})

	t.Run("RoleVal", func(t *testing.T) {
		input.Role = "magang"
		_, err := srv.Create(input)
		assert.ErrorContains(t, err, "role option only")
		data.AssertExpectations(t)
	})

	t.Run("Validator", func(t *testing.T) {
		input.Role = "user"
		input.FullName = ""
		_, err := srv.Create(input)
		assert.ErrorContains(t, err, "req")
		data.AssertExpectations(t)
	})

	t.Run("duplicated", func(t *testing.T) {
		input.FullName = "aaaa"
		data.On("Store", input).Return(uint(0), errors.New("duplicated")).Once()
		_, err := srv.Create(input)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	data := mocks.NewUserDataInterface(t)
	srv := New(data)
	var input = users.UserEntity{
		FullName: "aa",
		Email:    "mail@mail.com",
		Password: "123",
		Role:     "user",
	}
	t.Run("Success", func(t *testing.T) {
		data.On("SelectById", uint(1)).Return(mockData, nil)
		data.On("Edit", input, uint(1)).Return(uint(1), nil).Once()
		data.On("SelectById", uint(1)).Return(mockData, nil)
		resultData, err := srv.Update(uint(1), input)
		assert.NoError(t, err)
		assert.Equal(t, resultData.Email, input.Email)
		data.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		data.On("SelectById", uint(0)).Return(users.UserEntity{}, errors.New("not found"))
		_, err := srv.Update(uint(0), input)
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		data.On("SelectById", uint(1)).Return(mockData, nil)
		data.On("Edit", input, uint(1)).Return(uint(0), errors.New("not duplicated")).Once()
		_, err := srv.Update(uint(1), input)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewUserDataInterface(t)
	srv := New(data)

	t.Run("Success", func(t *testing.T) {
		data.On("SelectById", uint(1)).Return(mockData, nil)
		data.On("Destroy", uint(1)).Return(nil)
		err := srv.Delete(uint(1))
		assert.NoError(t, err)
		data.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		data.On("SelectById", uint(0)).Return(users.UserEntity{}, errors.New("not found"))
		err := srv.Delete(uint(0))
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})
}

func TestUpdateToSeller(t *testing.T) {
	data := mocks.NewUserDataInterface(t)
	srv := New(data)

	t.Run("Success", func(t *testing.T) {
		mockData.Address = "aa"
		mockData.PhoneNumber = "111"
		mockData.ShopName = "111"
		mockData.Role = "seller"
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		data.On("Edit", mockData, uint(1)).Return(uint(1), nil).Once()
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		_, err := srv.UpdateToSeller(uint(1), mockData)
		assert.NoError(t, err)
		data.AssertExpectations(t)
	})

	t.Run("Address", func(t *testing.T) {
		mockData.Address = ""
		mockData.PhoneNumber = "111"
		mockData.ShopName = "aa"
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		_, err := srv.UpdateToSeller(uint(1), mockData)
		assert.ErrorContains(t, err, "complete all your data first")
		data.AssertExpectations(t)
	})

	t.Run("RequestShopName", func(t *testing.T) {
		mockData.Address = "aa"
		mockData.PhoneNumber = "111"
		mockData.ShopName = ""
		mockData.Role = "seller"
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		_, err := srv.UpdateToSeller(uint(1), mockData)
		assert.ErrorContains(t, err, "insert shop name")
		data.AssertExpectations(t)
	})

	t.Run("RequestShopName", func(t *testing.T) {
		mockData.Address = "aa"
		mockData.PhoneNumber = "111"
		mockData.ShopName = "sas"
		mockData.Role = "seller"
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		data.On("Edit", mockData, uint(1)).Return(uint(1), errors.New("tes")).Once()
		_, err := srv.UpdateToSeller(uint(1), mockData)
		assert.NotNil(t, err)
		data.AssertExpectations(t)
	})

}
