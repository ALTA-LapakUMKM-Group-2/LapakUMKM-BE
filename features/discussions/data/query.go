package data

import (
	"lapakUmkm/features/discussions"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) discussions.DiscussionDataInterface {
	return &query{
		db: db,
	}
}

func (qd *query) Insert(discussionEntity discussions.DiscussionEntity) (uint, error) {
	discussion := DiscussionEntityToDiscussion(discussionEntity)
	if err := qd.db.Create(&discussion); err.Error != nil {
		return 0, err.Error
	}
	// //create discussion
	// if discussion.ParentId == 0 {
	// 	if err := qd.db.Create(&discussion); err.Error != nil {
	// 		return 0, err.Error
	// 	}
	// //create reply
	// } else {
	// 	if err := qd.db.Create(&discussion); err.Error != nil {
	// 		return 0, err.Error
	// 	}
	// }
	return discussion.ID, nil
}

func (qd *query) SelectById(id uint) (discussions.DiscussionEntity, error) {
	var discussion Discussion
	if err := qd.db.Preload("User").First(&discussion, id); err.Error != nil {
		return discussions.DiscussionEntity{}, err.Error
	}
	return DiscussionToDiscussionEntity(discussion), nil
}
