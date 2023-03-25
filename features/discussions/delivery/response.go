package delivery

import (
	"lapakUmkm/features/discussions"
	user "lapakUmkm/features/users/delivery"
	"reflect"
)

type DiscussionResponse struct {
	Id         uint   `json:"id"`
	ParentId   uint   `json:"parent_id" form:"parent_id"`
	UserId     uint   `json:"user_id" form:"user_id"`
	ProductId  uint   `json:"product_id" form:"product_id"`
	Discussion string `json:"discussion" form:"discussion"`
}

func DiscussionEntityToDiscussionResponse(discussionEntity discussions.DiscussionEntity) DiscussionResponse {
	return DiscussionResponse{
		Id:         discussionEntity.Id,
		ParentId:   discussionEntity.ParentId,
		UserId:     discussionEntity.UserId,
		ProductId:  discussionEntity.ProductId,
		Discussion: discussionEntity.Discussion,
	}
}

type DiscussionPutResponse struct {
	Id         uint   `json:"id"`
	Discussion string `json:"discussion" form:"discussion"`
}

func DiscussionEntityToDiscussionPutResponse(discussionEntity discussions.DiscussionEntity) DiscussionPutResponse {
	return DiscussionPutResponse{
		Id:         discussionEntity.Id,
		Discussion: discussionEntity.Discussion,
	}
}

type DiscussionGetResponse struct {
	Id         uint   `json:"id"`
	Discussion string `json:"discussion" form:"discussion"`
	User user.UserResponse `json:"user"`
}

func DiscussionEntityToDiscussionGetResponse(discussionEntity discussions.DiscussionEntity) DiscussionGetResponse {
	discussionGetResponse := DiscussionGetResponse{
		Id: discussionEntity.Id,
		Discussion: discussionEntity.Discussion,
	}
	if !reflect.ValueOf(discussionEntity.User).IsZero() {
		discussionGetResponse.User = user.UserResponse{
			FullName: discussionEntity.User.FullName,
			PhotoProfile: discussionEntity.User.PhotoProfile,
		}
	}
	return discussionGetResponse
}

func ListDiscussionToDiscussionGetResponse(discussionEntity []discussions.DiscussionEntity) []DiscussionGetResponse{
	var dataRes []DiscussionGetResponse
	var d discussions.DiscussionEntity 
	for _, v := range discussionEntity {
		if d.ParentId == 0 {
			dataRes = append(dataRes, DiscussionEntityToDiscussionGetResponse(v))
		}
	}
	return dataRes
}

func ListDiscussionEntityToDiscussionResponse(discussionEntity []discussions.DiscussionEntity) []DiscussionResponse {
	var dataResponse []DiscussionResponse
	for _, v := range discussionEntity {
		dataResponse = append(dataResponse, DiscussionEntityToDiscussionResponse(v))
	}
	return dataResponse
}
