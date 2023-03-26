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

func (qd *query) Store(discussionEntity discussions.DiscussionEntity) (uint, error) {
	discussion := DiscussionEntityToDiscussion(discussionEntity)
	if err := qd.db.Create(&discussion); err.Error != nil {
		return 0, err.Error
	}
	return discussion.ID, nil
}

func (qd *query) SelectById(id uint) (discussions.DiscussionEntity, error) {
	var discussion Discussion
	if err := qd.db.Preload("User").Preload("Product").First(&discussion, id); err.Error != nil {
		return discussions.DiscussionEntity{}, err.Error
	}
	return DiscussionToDiscussionEntity(discussion), nil
}

func (qd *query) Edit(discussionEntity discussions.DiscussionEntity, id uint) error {
	discussion := DiscussionEntityToDiscussion(discussionEntity)
	if err := qd.db.Where("id", id).Updates(&discussion); err.Error != nil {
		return err.Error
	}
	return nil
}

func (qd *query) Destroy(id uint) error {
	var discussion Discussion
	if err := qd.db.Delete(&discussion, id); err.Error != nil {
		return err.Error
	}
	return nil
}

func (qd *query) SelectDiscussionByProductId(productId uint) ([]discussions.DiscussionEntity, error) {
	discussion := []Discussion{}
	if err := qd.db.Where("product_id = ?", productId).Preload("User").Order("created_at desc").Find(&discussion).Error; err != nil {
		return []discussions.DiscussionEntity{}, err
	}
	res := []discussions.DiscussionEntity{}
	for _, v := range discussion {
		if v.ParentId == 0 {
			res = append(res, DiscussionToDiscussionEntity(v))
		}
	}
	return res, nil
}

func (qd *query) SelectAll(userId uint) ([]discussions.DiscussionEntity, error) {
	var discussions []Discussion
	if err := qd.db.Where("user_id = ?", userId).Preload("User").Order("created_at desc").Find(&discussions); err.Error != nil {
		return nil, err.Error
	}
	return ListDiscussionToDiscussionEntity(discussions), nil
}