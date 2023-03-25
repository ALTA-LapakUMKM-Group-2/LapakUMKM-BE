package data

import (
	"lapakUmkm/features/feedbacks"
	"lapakUmkm/features/products"
	product "lapakUmkm/features/products/data"
	"lapakUmkm/features/users"
	user "lapakUmkm/features/users/data"
	"reflect"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	ProductId uint
	Product   *product.Product `gorm:"foreignKey:ProductId"`
	ParentId  uint
	Rating    float64
	Feedback  string
	UserId    uint
	User      *user.User `gorm:"foreignKey:UserId"`
}

func FeedbackEntityToFeedback(feedbackEntity feedbacks.FeedbackEntity) Feedback {
	return Feedback{
		ProductId: feedbackEntity.ProductId,
		ParentId:  feedbackEntity.ParentId,
		Rating:    feedbackEntity.Rating,
		Feedback:  feedbackEntity.Feedback,
		UserId:    feedbackEntity.UserId,
	}
}

func FeedbackToFeedbackEntity(feedback Feedback) feedbacks.FeedbackEntity {
	result := feedbacks.FeedbackEntity{
		Id:        feedback.ID,
		ProductId: feedback.ProductId,
		ParentId:  feedback.ParentId,
		Rating:    feedback.Rating,
		Feedback:  feedback.Feedback,
		UserId:    feedback.UserId,
	}

	if !reflect.ValueOf(feedback.Product).IsZero() {
		result.Product = products.ProductEntity{
			Id: feedback.Product.ID,
		}
	}

	if !reflect.ValueOf(feedback.User).IsZero() {
		result.User = users.UserEntity{
			FullName:     feedback.User.FullName,
			PhotoProfile: feedback.User.PhotoProfile,
		}
	}
	return result
}

func ListFeedbackProductToFeedbackEntity(feedback []Feedback) []feedbacks.FeedbackEntity {
	var feedbackEntity []feedbacks.FeedbackEntity
	// var fb Feedback
	for _, v := range feedback {
		// if fb.ParentId == 0 {
			feedbackEntity = append(feedbackEntity, FeedbackToFeedbackEntity(v))
		// }
	}
	return feedbackEntity
}