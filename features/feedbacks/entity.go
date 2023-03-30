package feedbacks

import (
	"lapakUmkm/features/productTransactionDetails"
	"lapakUmkm/features/products"
	"lapakUmkm/features/users"
	"time"
)

type FeedbackEntity struct {
	Id                         uint
	ParentId                   uint
	UserId                     uint
	User                       users.UserEntity
	ProductId                  uint
	Product                    products.ProductEntity
	ProductTransactionDetailId uint
	ProductTransactionDetails  productTransactionDetails.ProductTransactionDetailEntity
	Rating                     float64
	Feedback                   string
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
}

type FeedbackServiceInterface interface {
	Create(feedbackEntity FeedbackEntity) (FeedbackEntity, error)
	Update(feedbackEntity FeedbackEntity, id, userId uint) (FeedbackEntity, error)
	Delete(id, userId uint) error
	GetFeedbackByProductId(productId uint) ([]FeedbackEntity, error)
	MyAllFeedback(myId, userId uint) ([]FeedbackEntity, error)
	GetById(id uint) (FeedbackEntity, error)
}

type FeedbackDataInterface interface {
	SelectById(id uint) (FeedbackEntity, error)
	Store(feedbackEntity FeedbackEntity) (uint, error)
	Edit(feedbackEntity FeedbackEntity, id uint) error
	Destroy(id uint) error
	SelectFeedbackByProductId(productId uint) ([]FeedbackEntity, error)
	SelectAll(userId uint) ([]FeedbackEntity, error)
}
