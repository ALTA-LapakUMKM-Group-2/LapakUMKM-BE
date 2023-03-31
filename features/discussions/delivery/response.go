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
	Childs     []Childs          `json:"childs,omitempty"`
}

type Childs struct {
	Id         uint              `json:"id"`
	ParentId   uint              `json:"parent_id" form:"parent_id"`
	UserId     uint              `json:"user_id" form:"user_id"`
	ProductId  uint              `json:"product_id" form:"product_id"`
	Discussion string            `json:"discussion" form:"discussion"`
	User       user.UserResponse `json:"user"`
}

func EntityToResponse(discussionEntity discussions.DiscussionEntity) DiscussionResponse {
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

func EntityToResponseChild(discussionEntity discussions.DiscussionEntity) Childs {
	discussionResponse := Childs{
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

func DiscussionEntityToDiscussionResponse(discussionEntity discussions.DiscussionEntity) DiscussionResponse {
	discussionResponse := EntityToResponse(discussionEntity)
	return discussionResponse
}

func ListDiscussionEntityToDiscussionResponse(discussionEntity []discussions.DiscussionEntity) []DiscussionResponse {
	var dataResponse []DiscussionResponse

	j := 0
	var flag uint
	if len(discussionEntity) > 0 {
		flag = discussionEntity[j].ParentId
	}

	for i := 0; i < len(discussionEntity); i++ {
		if flag != discussionEntity[i].ParentId {
			j++
			flag = discussionEntity[i].ParentId
		}

		if discussionEntity[i].Id == discussionEntity[i].ParentId {
			dataResponse = append(dataResponse, DiscussionEntityToDiscussionResponse(discussionEntity[i]))
		} else {
			dataResponse[j].Childs = append(dataResponse[j].Childs, EntityToResponseChild(discussionEntity[i]))
		}
	}

	return dataResponse
}

func AppendChilds(discussionEntity discussions.DiscussionEntity) DiscussionResponse {
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
