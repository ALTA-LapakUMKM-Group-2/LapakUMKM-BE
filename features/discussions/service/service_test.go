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

func TestCreate(t *testing.T) {
	repo := mocks.NewDiscussionDataInterface(t)
	srv := New(repo)
	inputData := discussions.DiscussionEntity{
		UserId:     uint(1),
		ProductId:  uint(1),
		ParentId:   0,
		Discussion: "ready?",
	}
	resData := discussions.DiscussionEntity{
		Id:         uint(1),
		ProductId:  uint(1),
		ParentId:   0,
		Discussion: "ready?",
		UserId:     uint(1),
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("Store", inputData).Return(uint(1), nil).Once()
		repo.On("SelectById", uint(1)).Return(resData, nil)
		resultData, err := srv.Create(inputData)
		assert.NoError(t, err)
		assert.Equal(t, resultData.Discussion, inputData.Discussion)
		repo.AssertExpectations(t)
	})

	t.Run("errVal", func(t *testing.T) {
		inputData.ProductId = uint(1)
		inputData.Discussion = ""
		_, err := srv.Create(inputData)
		assert.ErrorContains(t, err, "required")
		repo.AssertExpectations(t)
	})

	t.Run("errorDuplicated", func(t *testing.T) {
		expectedErr := errors.New("duplicated")
		repo.On("Store", mock.Anything).Return(uint(0), expectedErr).Once()
		res, err := srv.Create(resData)
		assert.Empty(t, res)
		assert.EqualError(t, err, expectedErr.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewDiscussionDataInterface(t)
	srv := New(repo)
	inputData := discussions.DiscussionEntity{
		UserId:     uint(1),
		ProductId:  uint(1),
		ParentId:   0,
		Discussion: "ready?",
	}
	resData := discussions.DiscussionEntity{
		Id:         uint(1),
		ProductId:  uint(1),
		ParentId:   0,
		Discussion: "ready?",
		UserId:     uint(1),
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("Edit", inputData, uint(1)).Return(nil).Once()
		repo.On("SelectById", uint(1)).Return(resData, nil).Twice()
		res, err := srv.Update(inputData, uint(1), uint(1))
		assert.NoError(t, err)
		assert.Equal(t, res.Discussion, inputData.Discussion)
		repo.AssertExpectations(t)
	})	

	// t.Run("errVal", func(t *testing.T) {
	// 	inputData.Discussion = ""
	// 	_, err := srv.Update(inputData, uint(1), uint(2))
	// 	assert.ErrorContains(t, err, "required")
	// 	repo.AssertExpectations(t)
	// })

	t.Run("notAuth", func(t *testing.T) {
        resData.UserId = 10
        repo.On("SelectById", uint(1)).Return(resData, nil).Once()
        _, err := srv.Update(inputData, uint(1), uint(2))
        assert.ErrorContains(t, err, "data can't be updated")
        repo.AssertExpectations(t)
    })
	
	t.Run("AccessDenied", func(t *testing.T) {
		repo.On("SelectById", uint(3)).Return(resData, nil).Once()
		_, err := srv.Update(inputData, uint(3), uint(2))
		assert.ErrorContains(t, err, "data can't be updated")
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewDiscussionDataInterface(t)
	srv := New(repo)
	resData := discussions.DiscussionEntity{
		Id:         uint(1),
		ProductId:  uint(1),
		ParentId:   0,
		Discussion: "ready?",
		UserId:     uint(1),
	}
	t.Run("Success", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
		repo.On("Destroy", uint(1)).Return(nil).Once()
		err := srv.Delete(uint(1), uint(1))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

    t.Run("notAuth", func(t *testing.T) {
        resData.UserId = 10
        repo.On("SelectById", uint(1)).Return(resData, nil).Once()
        err := srv.Delete(uint(1), uint(2))
        assert.ErrorContains(t, err, "access denied")
        repo.AssertExpectations(t)
    })
	
	t.Run("AccessDenied", func(t *testing.T) {
		repo.On("SelectById", uint(3)).Return(resData, nil).Once()
		err := srv.Delete(uint(3), uint(2))
		assert.ErrorContains(t, err, "access denied")
		repo.AssertExpectations(t)
	})
}

