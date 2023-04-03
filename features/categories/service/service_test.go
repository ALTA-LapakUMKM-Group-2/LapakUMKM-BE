package service

import (
	"errors"
	"lapakUmkm/features/categories"
	"lapakUmkm/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var inputData = categories.CategoryEntity{
	Category: "SEPATU",
}

var resData = categories.CategoryEntity{
	Id:       1,
	Category: "SEPATU",
}

var empty = categories.CategoryEntity{}

func TestGetAll(t *testing.T) {
	repo := mocks.NewCategoryDataInterface(t)
	dataExpected := []categories.CategoryEntity{
		{Id: 1, Category: "SEPATU"},
	}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(dataExpected, nil).Once()
		srv := New(repo)
		categories, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, dataExpected[0].Category, categories[0].Category)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := mocks.NewCategoryDataInterface(t)
	dataExpected := categories.CategoryEntity{
		Id: 1, Category: "SEPATU",
	}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(dataExpected, nil).Once()
		srv := New(repo)
		category, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, dataExpected.Category, category.Category)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := mocks.NewCategoryDataInterface(t)

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
		// repo.On("Store", mock.Anything).Return(uint(0), errors.New("required")).Once()
		srv := New(repo)
		_, err := srv.Create(categories.CategoryEntity{Category: ""})
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

func TestUpdate(t *testing.T) {
	repo := mocks.NewCategoryDataInterface(t)
	t.Run("errorValidation", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.Update(categories.CategoryEntity{Category: ""}, uint(1))
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "required")
		repo.AssertExpectations(t)
	})

	t.Run("notFound", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(empty, errors.New("not found")).Once()
		srv := New(repo)
		_, err := srv.Update(inputData, 1)
		assert.NotNil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("errorDuplicated", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
		repo.On("Edit", resData, uint(1)).Return(errors.New("duplicated")).Once()
		srv := New(repo)
		_, err := srv.Update(resData, uint(1))
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "duplicated")
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
		repo.On("Edit", resData, uint(1)).Return(nil).Once()
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()

		srv := New(repo)
		_, err := srv.Update(resData, uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewCategoryDataInterface(t)
	t.Run("success", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
		repo.On("Destroy", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("notFound", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(empty, errors.New("not found")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
		repo.On("Destroy", mock.Anything).Return(errors.New("error")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
}
