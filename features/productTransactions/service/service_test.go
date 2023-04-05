package service

import (
	"errors"
	"lapakUmkm/features/productTransactions"
	"lapakUmkm/features/users"
	"lapakUmkm/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewProductTransactionDataInterface(t)
	inputData := productTransactions.ProductTransactionEntity{
		UserId:       uint(1),
		TotalProduct: 30,
		TotalPayment: 30000,
	}

	t.Run("success", func(t *testing.T) {
		srv := New(repo)
		repo.On("Store", inputData).Return(uint(1), nil).Once()
		_, err := srv.Create(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("errVal", func(t *testing.T) {
		srv := New(repo)
		inputData.TotalProduct = 0
		_, err := srv.Create(inputData)
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("errDataStore", func(t *testing.T) {
		srv := New(repo)
		inputData.UserId = uint(0)
		inputData.TotalProduct = 2
		repo.On("Store", inputData).Return(uint(0), errors.New("duplicated")).Once()
		_, err := srv.Create(inputData)
		assert.ErrorContains(t, err, "duplicated")
		repo.AssertExpectations(t)
	})
}

func TestMyTransactionHistory(t *testing.T) {
	repo := mocks.NewProductTransactionDataInterface(t)
	resData := []productTransactions.ProductTransactionEntity{{
		Id:           1,
		UserId:       uint(1),
		TotalProduct: 30,
		TotalPayment: 304000,
	}, {
		Id:           2,
		UserId:       uint(1),
		TotalProduct: 304,
		TotalPayment: 440000,
	}}

	t.Run("success", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.MyTransactionHistory(1, 3)
		assert.Nil(t, err)
		assert.Equal(t, resData[1].Id, res[1].Id)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := mocks.NewProductTransactionDataInterface(t)
	resData := productTransactions.ProductTransactionEntity{
		Id:            1,
		UserId:        0,
		User:          users.UserEntity{},
		TotalProduct:  0,
		TotalPayment:  0,
		OrderId:       "",
		PaymentStatus: "",
		PaymentLink:   "",
	}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, resData.Id, feeds.Id)
		repo.AssertExpectations(t)
	})
}

func TestCallBackMidtrans(t *testing.T) {
	repo := mocks.NewProductTransactionDataInterface(t)
	inputData := productTransactions.ProductTransactionEntity{
		Id:            0,
		UserId:        0,
		TotalProduct:  0,
		TotalPayment:  0,
		OrderId:       "",
		PaymentStatus: "success",
		PaymentLink:   "",
	}
	status := "success"
	srv := New(repo)

	t.Run("success", func(t *testing.T) {
		repo.On("Edit", inputData, uint(1)).Return(nil).Once()
		err := srv.CallBackMidtrans(uint(1), status)

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("errorEdit", func(t *testing.T) {
		repo.On("Edit", inputData, uint(1)).Return(errors.New("error")).Once()
		err := srv.CallBackMidtrans(uint(1), status)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "error")
		repo.AssertExpectations(t)
	})
}
