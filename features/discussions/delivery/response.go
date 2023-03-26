package delivery

import (
	"lapakUmkm/features/discussions"
	user "lapakUmkm/features/users/delivery"
	"reflect"
)

type DiscussionResponse struct {
	Id         uint              `json:"id"`
	ParentId   uint              `json:"parent_id" form:"parent_id"`
	UserId     uint              `json:"user_id" form:"user_id"`
	ProductId  uint              `json:"product_id" form:"product_id"`
	Discussion string            `json:"discussion" form:"discussion"`
	User       user.UserResponse `json:"user"`
}

func DiscussionEntityToDiscussionResponse(discussionEntity discussions.DiscussionEntity) DiscussionResponse {
	discussionResponse := DiscussionResponse{
		Id:         discussionEntity.Id,
		ParentId:   discussionEntity.ParentId,
		UserId:     discussionEntity.UserId,
		ProductId:  discussionEntity.ProductId,
		Discussion: discussionEntity.Discussion,
	}
	if !reflect.ValueOf(discussionEntity.User).IsZero() {
		discussionResponse.User = user.UserResponse{
			FullName:     discussionEntity.User.FullName,
			PhotoProfile: discussionEntity.User.PhotoProfile,
		}
	}
	return discussionResponse
}

func ListDiscussionEntityToDiscussionResponse(discussionEntity []discussions.DiscussionEntity) []DiscussionResponse {
	var dataResponse []DiscussionResponse
	for _, v := range discussionEntity {
		dataResponse = append(dataResponse, DiscussionEntityToDiscussionResponse(v))
	}
	return dataResponse
}
