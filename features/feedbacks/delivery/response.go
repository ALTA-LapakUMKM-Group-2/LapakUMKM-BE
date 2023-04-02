package delivery

import (
	"lapakUmkm/features/feedbacks"

	// fb "lapakUmkm/features/feedbacks/data"
	"reflect"
	// product "lapakUmkm/features/products/delivery"
	user "lapakUmkm/features/users/delivery"
)

type FeedbackResponse struct {
	Id                         uint              `json:"id"`
	ParentId                   uint              `json:"parent_id"`
	ProductTransactionDetailId uint              `json:"product_transaction_detail_id"`
	ProductId                  uint              `json:"product_id"`
	Rating                     float64           `json:"rating"`
	Feedback                   string            `json:"feedback"`
	User                       user.UserResponse `json:"user"`
	Childs                     []Childs          `json:"childs,omitempty"`
}

type Childs struct {
	Id                         uint              `json:"id"`
	ParentId                   uint              `json:"parent_id"`
	ProductTransactionDetailId uint              `json:"product_transaction_detail_id"`
	ProductId                  uint              `json:"product_id"`
	Rating                     float64           `json:"rating"`
	Feedback                   string            `json:"feedback"`
	User                       user.UserResponse `json:"user"`
}

func FeedbackEntityToFeedbackResponse(feedbackEntity feedbacks.FeedbackEntity) FeedbackResponse {
	feedbackResponse := FeedbackResponse{
		Id:                         feedbackEntity.Id,
		ParentId:                   feedbackEntity.ParentId,
		ProductTransactionDetailId: feedbackEntity.ProductTransactionDetailId,
		ProductId:                  feedbackEntity.ProductId,
		Rating:                     feedbackEntity.Rating,
		Feedback:                   feedbackEntity.Feedback,
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

func EntityToResponseChild(f feedbacks.FeedbackEntity) Childs {
	discussionResponse := Childs{
		Id:                         f.Id,
		ParentId:                   f.ParentId,
		ProductId:                  f.ProductId,
		ProductTransactionDetailId: f.ProductTransactionDetailId,
		Rating:                     f.Rating,
		Feedback:                   f.Feedback,
	}
	if !reflect.ValueOf(f.User).IsZero() {
		discussionResponse.User = user.UserResponse{
			Id:           f.Id,
			FullName:     f.User.FullName,
			PhotoProfile: f.User.PhotoProfile,
		}
	}

	return discussionResponse
}

func AppendChilds(f feedbacks.FeedbackEntity) FeedbackResponse {
	discussionResponse := FeedbackResponse{
		Id:                         f.Id,
		ParentId:                   f.ParentId,
		ProductId:                  f.ProductId,
		ProductTransactionDetailId: f.ProductTransactionDetailId,
		Rating:                     f.Rating,
		Feedback:                   f.Feedback,
	}
	if !reflect.ValueOf(f.User).IsZero() {
		discussionResponse.User = user.UserResponse{
			Id:           f.Id,
			FullName:     f.User.FullName,
			PhotoProfile: f.User.PhotoProfile,
		}
	}

	return discussionResponse
}

func ListWithChild(f []feedbacks.FeedbackEntity) []FeedbackResponse {
	var dataResponse []FeedbackResponse
	j := 0
	var flag uint
	if len(f) > 0 {
		flag = f[j].ParentId
	}
	for i := 0; i < len(f); i++ {
		if flag != f[i].ParentId {
			j++
			flag = f[i].ParentId
		}

		if f[i].Id == f[i].ParentId {
			dataResponse = append(dataResponse, FeedbackEntityToFeedbackResponse(f[i]))
		} else {
			dataResponse[j].Childs = append(dataResponse[j].Childs, EntityToResponseChild(f[i]))
		}
	}

	return dataResponse
}
