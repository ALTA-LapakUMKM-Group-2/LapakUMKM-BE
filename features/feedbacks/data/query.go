package data

import (
	"lapakUmkm/features/feedbacks"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedbacks.FeedbackDataInterface{
	return &query{
		db: db,
	}
}

func (qf *query) Store(feedbackEntity feedbacks.FeedbackEntity) (uint, error) {
	feedback := FeedbackEntityToFeedback(feedbackEntity)
	if err := qf.db.Create(&feedback); err.Error != nil {
		return 0, err.Error
	}
	feedbackEntity.ProductId = feedback.ProductId
	return feedback.ID, nil
}

func (qf *query) SelectById(id uint) (feedbacks.FeedbackEntity, error) {
	var feedback Feedback
	if err := qf.db.Preload("User").Preload("Product").First(&feedback, id); err.Error != nil {
		return feedbacks.FeedbackEntity{}, err.Error
	}
	return FeedbackToFeedbackEntity(feedback), nil
}

func (qf *query) Edit(feedbackEntity feedbacks.FeedbackEntity, id uint) error {
	feedback := FeedbackEntityToFeedback(feedbackEntity)
	if err := qf.db.Where("id", id).Updates(&feedback); err.Error != nil {
		return err.Error
	}
	return nil
}

func (qf *query) Destroy(id uint) error{
	var feedback Feedback
	if err := qf.db.Delete(&feedback, id); err.Error != nil {
		return err.Error
	}
	return nil
}

func (qf *query) SelectFeedbackByProductId(detailTransactionId uint) ([]feedbacks.FeedbackEntity, error) {
	feedback := []Feedback{}
	if err := qf.db.Where("product_transaction_detail_id = ?", detailTransactionId).Preload("User").Order("created_at desc").Find(&feedback).Error; err != nil {
		return []feedbacks.FeedbackEntity{}, err
	}
	return ListFeedbackProductToFeedbackEntity(feedback), nil
}

func (qf *query) SelectFeedbackByDetailTransactionId(detailTransactionId uint) ([]feedbacks.FeedbackEntity, error) {
	feedback := []Feedback{}
	if err := qf.db.Where("product_id = ?", detailTransactionId).Preload("User").Order("created_at desc").Find(&feedback).Error; err != nil {
		return []feedbacks.FeedbackEntity{}, err
	}
	return ListFeedbackProductToFeedbackEntity(feedback), nil
}

func (qf *query) SelectAll(userId uint) ([]feedbacks.FeedbackEntity, error) {
	var feedback []Feedback
	if err := qf.db.Where("user_id = ?", userId).Preload("User").Order("created_at desc").Find(&feedback); err.Error != nil {
		return nil, err.Error
	}
	return ListFeedbackProductToFeedbackEntity(feedback), nil
}
