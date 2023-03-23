package service

import (
	"lapakUmkm/features/discussions"

	"github.com/go-playground/validator/v10"
)

type DiscussionService struct {
	Data discussions.DiscussionDataInterface
	validate *validator.Validate
}

func New(data discussions.DiscussionDataInterface) discussions.DiscussionServiceInterface{
	return &DiscussionService{
		Data: data,
		validate: validator.New(),
	}
}

func (sd *DiscussionService) Create(discussionEntity discussions.DiscussionEntity) (discussions.DiscussionEntity, error) {
	//validation
	sd.validate = validator.New()
	errValidate := sd.validate.StructExcept(discussionEntity, "User", "Product")
	if errValidate != nil {
		return discussions.DiscussionEntity{}, errValidate
	}
	//insertion
	user_id, err := sd.Data.Insert(discussionEntity)
	if err != nil {
		return discussions.DiscussionEntity{}, err 
	}
	return sd.Data.SelectById(user_id)
}