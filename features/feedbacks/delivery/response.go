package delivery

import (
	"lapakUmkm/features/feedbacks"
	// fb "lapakUmkm/features/feedbacks/data"
	"reflect"
	// product "lapakUmkm/features/products/delivery"
	user "lapakUmkm/features/users/delivery"
)

type FeedbackResponse struct {
	Id            uint              `json:"id"`
	ParentId      uint              `json:"parent_id"`
	TransactionId uint              `json:"transaction_id"`
	ProductId     uint              `json:"product_id"`
	Rating        float64           `json:"rating"`
	Feedback      string            `json:"feedback"`
	User          user.UserResponse `json:"user"`
}

func FeedbackEntityToFeedbackResponse(feedbackEntity feedbacks.FeedbackEntity) FeedbackResponse {
	feedbackResponse := FeedbackResponse{
		Id:            feedbackEntity.Id,
		ParentId:      feedbackEntity.ParentId,
		TransactionId: feedbackEntity.TransactionId,
		ProductId:     feedbackEntity.ProductId,
		Rating:        feedbackEntity.Rating,
		Feedback:      feedbackEntity.Feedback,
	}
	if !reflect.ValueOf(feedbackEntity.User).IsZero() {
		feedbackResponse.User = user.UserResponse{
			FullName:     feedbackEntity.User.FullName,
			PhotoProfile: feedbackEntity.User.PhotoProfile,
		}
	}
	return feedbackResponse
}

func ListFeedbackToFeedbackResponse(feedbackEntity []feedbacks.FeedbackEntity) []FeedbackResponse {
	var dataRes []FeedbackResponse
	for _, v := range feedbackEntity {
		dataRes = append(dataRes, FeedbackEntityToFeedbackResponse(v))
	}
	return dataRes
}


