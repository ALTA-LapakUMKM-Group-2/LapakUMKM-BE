package delivery

import "lapakUmkm/features/feedbacks"

type FeedbackRequest struct {
	ProductId                  uint    `json:"product_id" form:"product_id"`
	ParentId                   uint    `json:"parent_id" form:"parent_id"`
	Rating                     float64 `json:"rating" form:"rating"`
	Feedback                   string  `json:"feedback" form:"feedback"`
	ProductTransactionDetailId uint    `json:"product_transaction_detail_id" form:"product_transaction_detail_id"`
}

func FeedbackRequestToFeedbackEntity(feedbackRequest *FeedbackRequest) feedbacks.FeedbackEntity {
	return feedbacks.FeedbackEntity{
		ParentId:                   feedbackRequest.ParentId,
		ProductId:                  feedbackRequest.ProductId,
		Rating:                     feedbackRequest.Rating,
		Feedback:                   feedbackRequest.Feedback,
		ProductTransactionDetailId: feedbackRequest.ProductTransactionDetailId,
	}
}

type FeedbackPutRequest struct {
	Rating   float64 `json:"rating" form:"rating"`
	Feedback string  `json:"feedback" form:"feedback"`
}

func FeedbackPutRequestToFeedbackEntity(feedbackPutRequest *FeedbackPutRequest) feedbacks.FeedbackEntity {
	return feedbacks.FeedbackEntity{
		Rating:   feedbackPutRequest.Rating,
		Feedback: feedbackPutRequest.Feedback,
	}
}
