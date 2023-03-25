package delivery

import (
	"lapakUmkm/features/feedbacks"
	// fb "lapakUmkm/features/feedbacks/data"
	"reflect"
	// product "lapakUmkm/features/products/delivery"
	user "lapakUmkm/features/users/delivery"
	
)

type FeedbackPostResponse struct {
	Id        uint                    `json:"id"`
	ParentId uint `json:"parent_id"`
	UserId    uint                    `json:"user_id"`
	// User      user.UserResponse       `json:"user"`
	ProductId uint `json:"product_id"`
	// Product   product.ProductResponse `json:"product"`
	Rating   float64 `json:"rating"`
	Feedback string  `json:"feedback"`
}

func FeedbackEntityToFeedbackPostResponse(feedbackEntity feedbacks.FeedbackEntity) FeedbackPostResponse {
	feedbackPostResponse := FeedbackPostResponse{
		Id:        feedbackEntity.Id,
		ParentId: feedbackEntity.ParentId,
		UserId:    feedbackEntity.UserId,
		ProductId: feedbackEntity.ProductId,
		Rating:    feedbackEntity.Rating,
		Feedback:  feedbackEntity.Feedback,
	}
	return feedbackPostResponse
}

type FeedbackPutResponse struct {
	Id       uint    `json:"id"`
	Rating   float64 `json:"rating"`
	Feedback string  `json:"feedback"`
}

func FeedbackEntityToFeedbackPutResponse(feedbackEntity feedbacks.FeedbackEntity) FeedbackPutResponse {
	feedbackPutResponse := FeedbackPutResponse{
		Id:       feedbackEntity.Id,
		Rating:   feedbackEntity.Rating,
		Feedback: feedbackEntity.Feedback,
	}
	return feedbackPutResponse
}

type FeedbackGetResponse struct {
	Id       uint              `json:"id"`
	Rating   float64           `json:"rating"`
	Feedback string            `json:"feedback"`
	User     user.UserResponse `json:"user"`
}

func FeedbackEntityToFeedbackGetResponse(feedbackEntity feedbacks.FeedbackEntity) FeedbackGetResponse {
	feedbackGetResponse := FeedbackGetResponse{
		Id: feedbackEntity.Id,
		Rating: feedbackEntity.Rating,
		Feedback: feedbackEntity.Feedback,
	}
	if !reflect.ValueOf(feedbackEntity.User).IsZero() {
		feedbackGetResponse.User = user.UserResponse{
			FullName: feedbackEntity.User.FullName,
			PhotoProfile: feedbackEntity.User.PhotoProfile,
		}
	}
	return feedbackGetResponse
}

func ListFeedbackToFeedbackGetResponse(feedbackEntity []feedbacks.FeedbackEntity) []FeedbackGetResponse {
	var dataRes []FeedbackGetResponse
	var fb feedbacks.FeedbackEntity
	for _, v := range feedbackEntity {
		if fb.ParentId == 0 {
			dataRes = append(dataRes, FeedbackEntityToFeedbackGetResponse(v))
		}
	}
	return dataRes
}

func ListFeedbackGetAllToFeedbackGetResponse(feedbackEntity []feedbacks.FeedbackEntity) []FeedbackPostResponse {
	var dataRes []FeedbackPostResponse
	for _, v := range feedbackEntity {
			dataRes = append(dataRes, FeedbackEntityToFeedbackPostResponse(v))
	}
	return dataRes
}