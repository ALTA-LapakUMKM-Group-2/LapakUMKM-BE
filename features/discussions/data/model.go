package data

import (
	"lapakUmkm/features/discussions"
	"lapakUmkm/features/products"
	product "lapakUmkm/features/products/data"
	"lapakUmkm/features/users"
	user "lapakUmkm/features/users/data"
	"reflect"

	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	ProductId  uint
	Product    *product.Product `gorm:"foreignKey:ProductId"`
	UserId     uint
	User       *user.User `gorm:"foreignKey:UserId"`
	ParentId   uint       `gorm:"default:0"`
	Discussion string
}

func DiscussionEntityToDiscussion(discussionEntity discussions.DiscussionEntity) Discussion {
	return Discussion{
		ProductId:  discussionEntity.ProductId,
		ParentId:   discussionEntity.ParentId,
		Discussion: discussionEntity.Discussion,
		UserId:     discussionEntity.UserId,
	}
}

func DiscussionToDiscussionEntity(discussion Discussion) discussions.DiscussionEntity {
	result := discussions.DiscussionEntity{
		Id:         discussion.ID,
		ProductId:  discussion.ProductId,
		ParentId:   discussion.ParentId,
		UserId:     discussion.UserId,
		Discussion: discussion.Discussion,
		CreatedAt:  discussion.CreatedAt,
		UpdatedAt:  discussion.UpdatedAt,
	}
	if !reflect.ValueOf(discussion.Product).IsZero() {
		result.Product = products.ProductEntity{
			Id: discussion.Product.ID,
		}
	}

	if !reflect.ValueOf(discussion.User).IsZero() {
		result.User = users.UserEntity{
			FullName:     discussion.User.FullName,
			PhotoProfile: discussion.User.PhotoProfile,
		}
	}
	return result
}

func ListDiscussionToDiscussionEntity(discussion []Discussion) []discussions.DiscussionEntity {
	var discussionEntity []discussions.DiscussionEntity
	for _, v := range discussion {
		discussionEntity = append(discussionEntity, DiscussionToDiscussionEntity(v))
	}
	return discussionEntity
}
