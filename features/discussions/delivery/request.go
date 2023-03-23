package delivery

import (
	"lapakUmkm/features/discussions"

)

type DiscussionRequest struct {
	Discussion string `json:"discussion" form:"discussion"`
	ProductId uint `json:"product_id" form:"product_id"`
	ParentId uint `json:"parent_id" form:"parent_id"`
}

func DiscussionRequestToDiscussionEntity(discussionRequest *DiscussionRequest) discussions.DiscussionEntity {
	return discussions.DiscussionEntity{
		Discussion: discussionRequest.Discussion,
		ProductId: discussionRequest.ProductId,
		ParentId: discussionRequest.ParentId,
	}
}
