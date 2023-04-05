package service

import (
	"errors"
	"lapakUmkm/features/feedbacks"
	"lapakUmkm/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	inputData := feedbacks.FeedbackEntity{
		ParentId:                   1,
		UserId:                     1,
		ProductId:                  1,
		ProductTransactionDetailId: 1,
		Rating:                     4,
		Feedback:                   "mantap",
	}
	resData := feedbacks.FeedbackEntity{
		Id:                         1,
		ParentId:                   1,
		UserId:                     1,
		ProductId:                  1,
		ProductTransactionDetailId: 1,
		Rating:                     4,
		Feedback:                   "mantap",
	}

	t.Run("success", func(t *testing.T) {
		repo.On("Store", mock.Anything).Return(uint(1), nil).Once()
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()

		srv := New(repo)
		res, err := srv.Create(inputData)
		assert.NoError(t, err)
		assert.Equal(t, resData.Id, res.Id)
		repo.AssertExpectations(t)
	})

	t.Run("errorDuplicated", func(t *testing.T) {
		expectedErr := errors.New("duplicated")
		repo.On("Store", mock.Anything).Return(uint(0), expectedErr).Once()
		srv := New(repo)
		res, err := srv.Create(resData)
		assert.Empty(t, res)
		assert.EqualError(t, err, expectedErr.Error())
		repo.AssertExpectations(t)
	})

	t.Run("errorValidation", func(t *testing.T) {
		inputData.Feedback = ""
		srv := New(repo)
		_, err := srv.Create(inputData)
		assert.ErrorContains(t, err, "required")
		repo.AssertExpectations(t)
	})
}

func TestGetFeedbackByProductId(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	resData := []feedbacks.FeedbackEntity{{
		Id: 1, Rating: 4, Feedback: "sip",
	}, {
		Id: 2, Rating: 5, Feedback: "mantap",
	}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectFeedbackByProductId", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.GetFeedbackByProductId(1)
		assert.Nil(t, err)
		assert.Equal(t, resData[1].Feedback, feeds[1].Feedback)
		repo.AssertExpectations(t)
	})

	t.Run("errorInternalProblem", func(t *testing.T) {
		repo.On("SelectFeedbackByProductId", mock.Anything).Return(resData, errors.New("internal problem")).Once()
		srv := New(repo)
		_, err := srv.GetFeedbackByProductId(1)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "internal problem")
		repo.AssertExpectations(t)
	})
}

func TestGetFeedbackByDetailTransactionId(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	resData := []feedbacks.FeedbackEntity{{
		Id: 1, Rating: 4, Feedback: "sip",
	}, {
		Id: 2, Rating: 5, Feedback: "mantap",
	}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectFeedbackByDetailTransactionId", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.GetFeedbackByDetailTransactionId(1)
		assert.Nil(t, err)
		assert.Equal(t, resData[1].Feedback, feeds[1].Feedback)
		repo.AssertExpectations(t)
	})

	t.Run("errorInternalProblem", func(t *testing.T) {
		repo.On("SelectFeedbackByDetailTransactionId", mock.Anything).Return(resData, errors.New("internal problem")).Once()
		srv := New(repo)
		_, err := srv.GetFeedbackByDetailTransactionId(1)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "internal problem")
		repo.AssertExpectations(t)
	})
}

func TestMyAllFeedback(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	resData := []feedbacks.FeedbackEntity{{
		Id: 1, Rating: 4, Feedback: "sip",
	}, {
		Id: 2, Rating: 5, Feedback: "mantap",
	}}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.MyAllFeedback(1, 3)
		assert.Nil(t, err)
		assert.Equal(t, resData[1].Feedback, feeds[1].Feedback)
		repo.AssertExpectations(t)
	})

	t.Run("errorInternalProblem", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(resData, errors.New("internal problem")).Once()
		srv := New(repo)
		_, err := srv.MyAllFeedback(1, 3)
		assert.NotEmpty(t, err)
		assert.ErrorContains(t, err, "internal problem")
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	resData := feedbacks.FeedbackEntity{
		Id: 1, Feedback: "oke",
	}

	t.Run("-", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		feeds, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, resData.Feedback, feeds.Feedback)
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

func TestUpdate(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	srv := New(repo)
	inputData := feedbacks.FeedbackEntity{
		ProductId: uint(1),
		UserId:    uint(1),
		ParentId:  0,
		Rating:    3,
		Feedback:  "bagus",
	}
	resData := feedbacks.FeedbackEntity{
		Id:        uint(1),
		ProductId: uint(1),
		UserId:    uint(1),
		ParentId:  0,
		Rating:    3,
		Feedback:  "bagus",
	}

	t.Run("Success", func(t *testing.T) {
		repo.On("Edit", inputData, uint(1)).Return(nil).Once()
		repo.On("SelectById", uint(1)).Return(resData, nil).Twice()
		res, err := srv.Update(inputData, uint(1), uint(1))
		assert.NoError(t, err)
		assert.Equal(t, res.Feedback, inputData.Feedback)
		repo.AssertExpectations(t)
	})
	t.Run("notAuth", func(t *testing.T) {
		resData.UserId = 10
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
		_, err := srv.Update(inputData, uint(1), uint(2))
		assert.ErrorContains(t, err, "access denied")
		repo.AssertExpectations(t)
	})

	t.Run("AccessDenied", func(t *testing.T) {
		repo.On("SelectById", uint(3)).Return(resData, nil).Once()
		_, err := srv.Update(inputData, uint(3), uint(2))
		assert.ErrorContains(t, err, "access denied")
		repo.AssertExpectations(t)
	})

	t.Run("notfound", func(t *testing.T) {
		repo.On("SelectById", uint(31)).Return(feedbacks.FeedbackEntity{}, errors.New("not found")).Once()
		_, err := srv.Update(inputData, uint(31), uint(2))
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("ErrorEdit", func(t *testing.T) {
		inputData.Feedback = ""
		repo.On("Edit", inputData, uint(1)).Return(errors.New("required")).Once()
		repo.On("SelectById", uint(1)).Return(inputData, nil).Once()
		_, err := srv.Update(inputData, uint(1), uint(1))
		assert.ErrorContains(t, err, "required")
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	srv := New(repo)
	resData := feedbacks.FeedbackEntity{
		Id:        uint(1),
		ProductId: uint(1),
		UserId:    uint(1),
		ParentId:  0,
		Rating:    3,
		Feedback:  "bagus",
	}
	t.Run("Success", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
		repo.On("Destroy", uint(1)).Return(nil).Once()
		err := srv.Delete(uint(1), uint(1))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("notFound", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(feedbacks.FeedbackEntity{}, errors.New("not found")).Once()
		srv := New(repo)
		err := srv.Delete(uint(1), uint(2))
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
