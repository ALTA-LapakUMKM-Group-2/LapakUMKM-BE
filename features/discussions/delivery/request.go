package delivery

import (
	"lapakUmkm/features/discussions"
)

type DiscussionRequest struct {
	ProductId  uint   `json:"product_id" form:"product_id"`
	ParentId   uint   `json:"parent_id" form:"parent_id"`
	Discussion string `json:"discussion" form:"discussion"`
}

func DiscussionRequestToDiscussionEntity(discussionRequest *DiscussionRequest) discussions.DiscussionEntity {
	return discussions.DiscussionEntity{
		ProductId:  discussionRequest.ProductId,
		ParentId:   discussionRequest.ParentId,
		Discussion: discussionRequest.Discussion,
	}
}

type DiscussionPutRequest struct {
	Discussion string `json:"discussion" form:"discussion"`
}

func DiscussionPutRequestToDiscussionEntity(discussionPutRequest *DiscussionPutRequest) discussions.DiscussionEntity {
	return discussions.DiscussionEntity{
		Discussion: discussionPutRequest.Discussion,
	}
}