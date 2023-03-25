package discussions

import (
	"lapakUmkm/features/products"
	"lapakUmkm/features/users"
	"time"
)

type DiscussionEntity struct {
	Id         uint
	ProductId  uint
	Product    products.ProductEntity
	ParentId   uint
	Discussion string `validate:"required"`
	UserId     uint
	User       users.UserEntity
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DiscussionServiceInterface interface {
	Create(discussionEntity DiscussionEntity) (DiscussionEntity, error)
	Update(discussionEntity DiscussionEntity, id, userId uint) (DiscussionEntity, error)
	Delete(id, userId uint) error
	// GetFeedbackByProductId(productId uint) ([]DiscussionEntity, error)
	// GetAll() ([]DiscussionEntity, error)
	// GetById(id uint) (DiscussionEntity, error)
}

type DiscussionDataInterface interface {
	SelectById(id uint) (DiscussionEntity, error)
	Store(discussionEntity DiscussionEntity) (uint, error)
	Edit(discussionEntity DiscussionEntity, id uint) error
	Destroy(id uint) error
	// SelectFeedbackByProductId(productId uint) ([]DiscussionEntity, error)
	// SelectAll() ([]DiscussionEntity, error)

}
