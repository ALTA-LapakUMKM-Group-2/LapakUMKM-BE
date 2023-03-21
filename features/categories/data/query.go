package data

import (
	"lapakUmkm/features/categories"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) categories.CategoryDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]categories.CategoryEntity, error) {
	var categories []Category
	if err := q.db.Preload("User").Find(&categories); err.Error != nil {
		return nil, err.Error
	}
	return ListCategoryToCategoryEntity(categories), nil
}

func (q *query) SelectById(id uint) (categories.CategoryEntity, error) {
	var category Category
	if err := q.db.Preload("User").First(&category, id); err.Error != nil {
		return categories.CategoryEntity{}, err.Error
	}
	return CategoryToCategoryEntity(category), nil
}

func (q *query) Store(categoryEntity categories.CategoryEntity) (uint, error) {
	category := CategoryEntityToCategory(categoryEntity)
	if err := q.db.Create(&category); err.Error != nil {
		return 0, err.Error
	}
	return category.ID, nil
}

func (q *query) Edit(categoryEntity categories.CategoryEntity, id uint) error {
	category := CategoryEntityToCategory(categoryEntity)
	if err := q.db.Where("id", id).Updates(&category); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var category Category
	if err := q.db.Delete(&category, id); err.Error != nil {
		return err.Error
	}
	return nil
}
