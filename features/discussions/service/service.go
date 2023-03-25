package service

import (
	// "errors"
	"lapakUmkm/features/discussions"
	// "log"

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
	errValidate := sd.validate.StructExcept(discussionEntity, "Product", "User")
	if errValidate != nil {
		return discussions.DiscussionEntity{}, errValidate
	}
	//insertion
	discussionId, err := sd.Data.Store(discussionEntity)
	if err != nil {
		return discussions.DiscussionEntity{}, err 
	}
	return sd.Data.SelectById(discussionId)
}

func (sd *DiscussionService) Update(discussionEntity discussions.DiscussionEntity, id uint, userId uint) (discussions.DiscussionEntity, error) {
	checkDataExist, errData := sd.Data.SelectById(id)
	if errData != nil {
		return checkDataExist, errData
	}

	if checkDataExist.UserId != userId {
		return discussions.DiscussionEntity{} , errors.New("data can't be updated")
	}

	err := sd.Data.Edit(discussionEntity, id)
	if err != nil {
		return discussions.DiscussionEntity{}, err
	}
	return sd.Data.SelectById(id)
}