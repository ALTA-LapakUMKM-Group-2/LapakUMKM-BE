package service

import (
	"errors"
	"lapakUmkm/features/products"
	"lapakUmkm/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockData = products.ProductEntity{
	Id:             uint(1),
	UserId:         uint(1),
	CategoryId:     uint(1),
	ProductName:    "aaa",
	Description:    "bbb",
	Price:          100,
	StockRemaining: 10,
	Size:           "L",
	StockSold:      10,
}

func TestGetAll(t *testing.T) {
	data := mocks.NewProductDataInterface(t)
	srv := New(data)
	listData := []products.ProductEntity{}
	listData = append(listData, mockData)

	filter := products.ProductFilter{}

	t.Run("Success", func(t *testing.T) {
		data.On("SelectAll", filter).Return(listData, nil).Once()
		_, err := srv.GetAll(filter)
		assert.NoError(t, err)
		data.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	data := mocks.NewProductDataInterface(t)
	srv := New(data)

	t.Run("Success", func(t *testing.T) {
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		_, err := srv.GetById(uint(1))
		assert.NoError(t, err)
		data.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	data := mocks.NewProductDataInterface(t)
	srv := New(data)
	var input = products.ProductEntity{
		UserId:         uint(1),
		CategoryId:     uint(1),
		ProductName:    "aaa",
		Description:    "bbb",
		Price:          100,
		StockRemaining: 10,
		Size:           "L",
		StockSold:      10,
	}

	t.Run("Success", func(t *testing.T) {
		data.On("Store", input).Return(uint(1), nil).Once()
		data.On("SelectById", uint(1)).Return(mockData, nil)
		resultData, err := srv.Create(input)
		assert.NoError(t, err)
		assert.Equal(t, resultData.ProductName, input.ProductName)
		data.AssertExpectations(t)
	})

	t.Run("RoleVal", func(t *testing.T) {
		input.ProductName = ""
		_, err := srv.Create(input)
		assert.ErrorContains(t, err, "required")
		data.AssertExpectations(t)
	})

	t.Run("duplicated", func(t *testing.T) {
		input.ProductName = "aaaa"
		data.On("Store", input).Return(uint(0), errors.New("duplicated")).Once()
		_, err := srv.Create(input)
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	data := mocks.NewProductDataInterface(t)
	srv := New(data)
	var input = products.ProductEntity{
		UserId:         uint(1),
		CategoryId:     uint(1),
		ProductName:    "aaa",
		Description:    "bbb",
		Price:          100,
		StockRemaining: 10,
		Size:           "L",
		StockSold:      10,
	}

	t.Run("Success", func(t *testing.T) {
		data.On("Edit", input, uint(1)).Return(nil).Once()
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		resultData, err := srv.Update(input, uint(1), uint(1))
		assert.NoError(t, err)
		assert.Equal(t, resultData.ProductName, input.ProductName)
		data.AssertExpectations(t)
	})

	t.Run("RoleVal", func(t *testing.T) {
		input.ProductName = ""
		_, err := srv.Update(input, uint(1), uint(1))
		assert.ErrorContains(t, err, "required")
		data.AssertExpectations(t)
	})

	t.Run("duplicated", func(t *testing.T) {
		input.ProductName = "aaaa"
		data.On("Edit", input, uint(1)).Return(errors.New("duplicated")).Once()
		_, err := srv.Update(input, uint(1), uint(1))
		assert.ErrorContains(t, err, "duplicated")
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewProductDataInterface(t)
	srv := New(data)

	t.Run("Success", func(t *testing.T) {
		data.On("SelectById", uint(1)).Return(mockData, nil).Once()
		data.On("Destroy", uint(1)).Return(nil).Once()
		err := srv.Delete(uint(1), uint(1))
		assert.NoError(t, err)
		data.AssertExpectations(t)
	})

	t.Run("notFound", func(t *testing.T) {
		data.On("SelectById", uint(2)).Return(products.ProductEntity{}, errors.New("not found")).Once()
		err := srv.Delete(uint(2), uint(1))
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})

	t.Run("notAuth", func(t *testing.T) {
		mockData.UserId = 10
		data.On("SelectById", uint(1)).Return(mockData, errors.New("don't have access to delete this product id")).Once()
		err := srv.Delete(uint(1), uint(2))
		assert.ErrorContains(t, err, "don't have access to delete this product id")
		data.AssertExpectations(t)
	})

	t.Run("noAccess", func(t *testing.T) {
		data.On("SelectById", uint(3)).Return(mockData, nil).Once()
		err := srv.Delete(uint(3), uint(2))
		assert.ErrorContains(t, err, "don't have access to delete this product id")
		data.AssertExpectations(t)
	})
}
