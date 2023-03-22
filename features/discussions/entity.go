package discussions

import (
	"lapakUmkm/features/products"
	"time"
)

type DiscussionEntity struct {
	Id         uint
	ProductId  uint
	Product    products.ProductEntity
	ParentID   uint
	Discussion string `validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DiscussionServiceInterface interface {
	Create(discussionEntity DiscussionEntity) (DiscussionEntity, error)
}

type DiscussionDataInterface interface {
	Insert(discussionEntity DiscussionEntity) (uint, error)
	SelectById(id uint) (DiscussionEntity, error)

}