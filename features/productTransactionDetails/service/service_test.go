package service

import (
	"errors"
	"lapakUmkm/features/productTransactionDetails"
	"lapakUmkm/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	repo := mocks.NewProductTransactionDetailDataInterface(t)
	// inputData :=productTransactionDetails.ProductTransactionDetailEntity{}
	resData := productTransactionDetails.ProductTransactionDetailEntity{}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		data, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, resData.Id, data.Id)
		repo.AssertExpectations(t)
	})
}

func TestGetByTransaksiId(t *testing.T) {
	repo := mocks.NewProductTransactionDetailDataInterface(t)
	// inputData :=productTransactionDetails.ProductTransactionDetailEntity{}
	resData := []productTransactionDetails.ProductTransactionDetailEntity{{},{}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectByTransaksiId", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		data, err := srv.GetByTransaksiId(1)
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ProductId, data[0].ProductId)
		repo.AssertExpectations(t)
	})
}


func TestCreate(t *testing.T) {
	repo := mocks.NewProductTransactionDetailDataInterface(t)
	inputData :=productTransactionDetails.ProductTransactionDetailEntity{}
	resData := productTransactionDetails.ProductTransactionDetailEntity{}

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
		assert.NotNil(t, err)
		// assert.NotEmpty(t, err)
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
