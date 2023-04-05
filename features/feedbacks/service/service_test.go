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
	inputData := feedbacks.FeedbackEntity{}
	resData := feedbacks.FeedbackEntity{}

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
		expectedErr := errors.New("required")
		repo.On("Store", mock.Anything).Return(uint(0), expectedErr).Once()
		srv := New(repo)
		res, err := srv.Create(inputData)
		assert.Empty(t, res)
		assert.Equal(t, feedbacks.FeedbackEntity{}, res)
		assert.EqualError(t, err, expectedErr.Error())
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

// func TestUpdate(t *testing.T) {
// 	repo := mocks.NewFeedbackDataInterface(t)
// 	inputData := feedbacks.FeedbackEntity{
// 		Rating: 4, Feedback: "oke",
// 	}
// 	resData := feedbacks.FeedbackEntity{
// 		Id: 1, Rating: 4, Feedback: "oke",
// 	}

// 	t.Run("errorValidation", func(t *testing.T) {
// 		expectedErr := errors.New("required")
// 		repo.On("Edit", mock.Anything).Return(uint(0), expectedErr).Once()
// 		repo.On("SelectById", mock.Anything).Return(uint(0), expectedErr).Once()
// 		srv := New(repo)
// 		_, err := srv.Update(feedbacks.FeedbackEntity{}, uint(1), uint(1))
// 		// assert.Empty(t, res)
// 		// assert.Equal(t, feedbacks.FeedbackEntity{}, res)
// 		// assert.EqualError(t, err, expectedErr.Error())
// 		// repo.AssertExpectations(t)
// 		assert.NotEmpty(t, err)
// 		assert.ErrorContains(t, err, "required")
// 		repo.AssertExpectations(t)
// 	})


// 	t.Run("notFound", func(t *testing.T) {
// 		repo.On("SelectById", uint(1)).Return(feedbacks.FeedbackEntity{}, errors.New("not found")).Once()
// 		srv := New(repo)
// 		_, err := srv.Update(inputData, 1, 1)
// 		assert.NotNil(t, err)
// 		assert.NotEmpty(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("errorDuplicated", func(t *testing.T) {
// 		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
// 		repo.On("Edit", resData, uint(1)).Return(errors.New("duplicated")).Once()
// 		srv := New(repo)
// 		_, err := srv.Update(resData, uint(1), uint(1))
// 		assert.NotEmpty(t, err)
// 		assert.ErrorContains(t, err, "duplicated")
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("success", func(t *testing.T) {
// 		repo.On("SelectById", uint(1)).Return(resData, nil).Once()
// 		repo.On("Edit", resData, uint(1)).Return(nil).Once()
// 		// repo.On("SelectById", uint(1)).Return(resData, nil).Once()

// 		srv := New(repo)
// 		_, err := srv.Update(resData, uint(1), uint(2))
// 		assert.NotEqual(t, resData, err)
// 		assert.Nil(t, nil)
// 		repo.AssertExpectations(t)
// 	})
// }

