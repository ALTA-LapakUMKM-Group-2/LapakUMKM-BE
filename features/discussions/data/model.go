package data

import (
	"lapakUmkm/features/discussions"
	"lapakUmkm/features/products/data"

	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	ProductId  uint
	Product    *data.Product `gorm:"foreignKey:ProductId"`
	ParentId   uint          `gorm:"default:0"`
	Discussion string
}

func DiscussionEntityToDiscussion(discussionEntity discussions.DiscussionEntity) Discussion {
	return Discussion{
		Discussion: discussionEntity.Discussion,
	}
}

func DiscussionToDiscussionEntity(discussion Discussion) discussions.DiscussionEntity {
	result := discussions.DiscussionEntity{
		Id:         discussion.ID,
		ProductId:  discussion.ProductId,
		ParentId:   discussion.ParentId,
		Discussion: discussion.Discussion,
		CreatedAt:  discussion.CreatedAt,
		UpdatedAt:  discussion.UpdatedAt,
	}
	return result
}
