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