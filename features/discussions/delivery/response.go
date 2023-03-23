package delivery

import "lapakUmkm/features/discussions"

type DiscussionResponse struct {
	Id         uint   `json:"id"`
	Discussion string `json:"discussion" form:"discussion"`
	ProductId  uint   `json:"product_id" form:"product_id"`
	ParentId   uint   `json:"parent_id" form:"parent_id"`
}

func DiscussionEntityToDiscussionResponse(discussionEntity discussions.DiscussionEntity) DiscussionResponse {
	return DiscussionResponse{
		Id:         discussionEntity.Id,
		Discussion: discussionEntity.Discussion,
		ProductId:  discussionEntity.ProductId,
		ParentId:   discussionEntity.ParentId,
	}
}