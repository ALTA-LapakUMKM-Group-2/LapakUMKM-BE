package service

import (
	"errors"
	"lapakUmkm/features/feedbacks"
	"log"

	// fb "lapakUmkm/features/feedbacks/data"
	"github.com/go-playground/validator/v10"
)

type feedbackService struct {
	Data feedbacks.FeedbackDataInterface
	validate *validator.Validate
}

func New(data feedbacks.FeedbackDataInterface) feedbacks.FeedbackServiceInterface {
	return &feedbackService{
		Data: data,
		validate: validator.New(),
	}
}

func (sf *feedbackService) Create(feedbackEntity feedbacks.FeedbackEntity) (feedbacks.FeedbackEntity, error) {
	sf.validate = validator.New()
	errValidate := sf.validate.StructExcept(feedbackEntity, "User", "Product")
	if errValidate != nil {
		return feedbacks.FeedbackEntity{}, errValidate
	}
	feedbackId, err := sf.Data.Store(feedbackEntity)
	if err != nil {
		return feedbacks.FeedbackEntity{}, err
	}
	return sf.Data.SelectById(feedbackId)
}

func (sf *feedbackService) Update(feedbackEntity feedbacks.FeedbackEntity, id, userId uint) (feedbacks.FeedbackEntity, error) {
	checkDataExist, errData := sf.Data.SelectById(id)
	if errData != nil {
		return checkDataExist, errData
	}

	if checkDataExist.UserId != userId {
		return feedbacks.FeedbackEntity{} , errors.New("data can't be updated")
	}

	err := sf.Data.Edit(feedbackEntity, id)
	if err != nil {
		return feedbacks.FeedbackEntity{}, err
	}
	return sf.Data.SelectById(id)
}