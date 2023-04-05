package service

import (
	"errors"
	"lapakUmkm/features/carts"
	"lapakUmkm/mocks"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCartService_Add(t *testing.T) {
	mockRepo := new(mocks.CartData)
	srv := New(mockRepo)

	newCart := carts.Core{
		UserId:     1,
		ProductId:  2,
		ProductPcs: 3,
	}

	expectedCart := carts.Core{
		Id:         1,
		ProductId:  newCart.ProductId,
		ProductPcs: newCart.ProductPcs,
	}

	mockRepo.On("Add", mock.Anything).Return(expectedCart, nil).Once()

	actualCart, err := srv.Add(newCart)
	assert.NoError(t, err)
	assert.Equal(t, expectedCart, actualCart)
	mockRepo.AssertExpectations(t)

	mockRepo.On("Add", newCart).Return(carts.Core{}, errors.New("failed add product to cart")).Once()

	_, err = srv.Add(newCart)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)

	// Test case 2: validation error when adding a new cart
	t.Run("errorValidation", func(t *testing.T) {
		srv := New(mockRepo)
		_, err := srv.Add(carts.Core{UserId: 3, ProductId: 0, ProductPcs: 2})
		assert.NotEmpty(t, err)
		assert.False(t, strings.Contains(err.Error(), "validation error"))
		mockRepo.AssertExpectations(t)
	})
}

func TestCartService_MyCart(t *testing.T) {
	mockRepo := new(mocks.CartData)
	srv := New(mockRepo)

	expectedCart := []carts.Core{
		{Id: 1, UserId: 1, ProductId: 1, ProductName: "Product A", ProductPrice: 10000, ProductPcs: 3, ProductImage: "image1.jpg", LapakName: "Lapak A", LapakAddress: "Indonesia", PhotoProfile: "profile1"},
		{Id: 2, UserId: 1, ProductId: 2, ProductName: "Product B", ProductPrice: 20000, ProductPcs: 2, ProductImage: "image2.jpg", LapakName: "Lapak B", LapakAddress: "Indonesia", PhotoProfile: "profile2"},
	}
	mockRepo.On("MyCart", uint(1)).Return(expectedCart, nil).Once()

	actualCart, err := srv.MyCart(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedCart, actualCart)
	mockRepo.AssertExpectations(t)

	mockRepo.On("MyCart", uint(1)).Return(nil, errors.New("failed show mycart")).Once()

	_, err = srv.MyCart(1)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}


func TestCartService_Update(t *testing.T) {
	mockRepo := new(mocks.CartData)
	srv := New(mockRepo)

	updateCart := carts.Core{
		Id:         1,
		ProductPcs: 5,
	}
	expectedCart := updateCart

	mockRepo.On("Update", updateCart).Return(expectedCart, nil).Once()

	actualCart, err := srv.Update(updateCart)
	assert.NoError(t, err)
	assert.Equal(t, expectedCart, actualCart)
	mockRepo.AssertExpectations(t)

	mockRepo.On("Update", updateCart).Return(carts.Core{}, errors.New("failed to update cart")).Once()

	_, err = srv.Update(updateCart)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCartService_Delete(t *testing.T) {
	mockRepo := new(mocks.CartData)
	srv := New(mockRepo)

	userID := uint(1)
	cartID := uint(2)

	mockRepo.On("Delete", userID, cartID).Return(nil).Once()

	err := srv.Delete(userID, cartID)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	
	mockRepo.On("Delete", userID, cartID).Return(errors.New("failed delete cart")).Once()

	err = srv.Delete(1, 2)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
