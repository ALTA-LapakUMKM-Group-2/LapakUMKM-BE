package service

import (
	"errors"
	"lapakUmkm/features/feedbacks"
	"log"

	"github.com/go-playground/validator/v10"
)

type feedbackService struct {
	Data     feedbacks.FeedbackDataInterface
	validate *validator.Validate
}

func New(data feedbacks.FeedbackDataInterface) feedbacks.FeedbackServiceInterface {
	return &feedbackService{
		Data:     data,
		validate: validator.New(),
	}
}

func (sf *feedbackService) Create(feedbackEntity feedbacks.FeedbackEntity) (feedbacks.FeedbackEntity, error) {
	sf.validate = validator.New()
	errValidate := sf.validate.StructExcept(feedbackEntity, "User", "Product", "ProductTransactionDetails")
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
		return feedbacks.FeedbackEntity{}, errors.New("access denied")
	}

	err := sf.Data.Edit(feedbackEntity, id)
	if err != nil {
		return feedbacks.FeedbackEntity{}, err
	}
	return sf.Data.SelectById(id)
}

func (sf *feedbackService) Delete(id, userId uint) error {
	checkDataExist, err := sf.Data.SelectById(id)
	if err != nil {
		return nil
	}
	if checkDataExist.UserId != userId {
		return errors.New("access denied")
	}
	return sf.Data.Destroy(id)
}

func (sf *feedbackService) GetFeedbackByProductId(productId uint) ([]feedbacks.FeedbackEntity, error) {
	res, err := sf.Data.SelectFeedbackByProductId(productId)
	if err != nil {
		log.Println("query error", err.Error())
		return []feedbacks.FeedbackEntity{}, errors.New("internal problem")
	}
	return res, nil
}

func (sf *feedbackService) GetFeedbackByDetailTransactionId(detailTransactionId uint) ([]feedbacks.FeedbackEntity, error) {
	res, err := sf.Data.SelectFeedbackByDetailTransactionId(detailTransactionId)
	if err != nil {
		log.Println("query error", err.Error())
		return []feedbacks.FeedbackEntity{}, errors.New("internal problem")
	}
	return res, nil
}


func (sf *feedbackService) MyAllFeedback(myId, userId uint) ([]feedbacks.FeedbackEntity, error) {
	return sf.Data.SelectAll(userId)
}

func (sf *feedbackService) GetById(id uint) (feedbacks.FeedbackEntity, error) {
	return sf.Data.SelectById(id)
}
