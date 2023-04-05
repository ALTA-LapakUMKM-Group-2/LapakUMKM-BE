package service

import (
	"errors"
	"lapakUmkm/features/discussions"
	"lapakUmkm/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDiscussionByProductId(t *testing.T) {
	repo := mocks.NewDiscussionDataInterface(t)
	resData := []discussions.DiscussionEntity{{
		Id: 1, Discussion: "",
	}, {
		Id: 2, Discussion: "",
	}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectDiscussionByProductId", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.GetDiscussionByProductId(1)
		assert.Nil(t, err)
		assert.Equal(t, resData[1].Discussion, feeds[1].Discussion)
		repo.AssertExpectations(t)
	})

	t.Run("errorInternalProblem", func(t *testing.T) {
		repo.On("SelectDiscussionByProductId", mock.Anything).Return(resData, errors.New("internal problem")).Once()
		srv := New(repo)
		_, err := srv.GetDiscussionByProductId(1)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "internal problem")
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := mocks.NewDiscussionDataInterface(t)
	resData := []discussions.DiscussionEntity{{
		Id: 1, Discussion: "",
	}, {
		Id: 2, Discussion: "",
	}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.GetAll(1, 3)
		assert.Nil(t, err)
		assert.Equal(t, resData[1].Discussion, feeds[1].Discussion)
		repo.AssertExpectations(t)
	})

	t.Run("errorInternalProblem", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(resData, errors.New("internal problem")).Once()
		srv := New(repo)
		_, err := srv.GetAll(1, 3)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "internal problem")
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := mocks.NewDiscussionDataInterface(t)
	resData := discussions.DiscussionEntity{
		Id: 1, Discussion: "",
	}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, resData.Discussion, feeds.Discussion)
		repo.AssertExpectations(t)
	})

	t.Run("errorInternalProblem", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(resData, errors.New("internal problem")).Once()
		srv := New(repo)
		_, err := srv.GetById(1)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "internal problem")
		repo.AssertExpectations(t)
	})
}