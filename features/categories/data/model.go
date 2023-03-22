package data

import (
	"lapakUmkm/features/categories"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category string
}

func CategoryEntityToCategory(categoryEntity categories.CategoryEntity) Category {
	return Category{
		Category: categoryEntity.Category,
	}
}

func CategoryToCategoryEntity(category Category) categories.CategoryEntity {
	result := categories.CategoryEntity{
		Id:        category.ID,
		Category:  category.Category,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
	return result
}

func ListCategoryToCategoryEntity(category []Category) []categories.CategoryEntity {
	var categoryEntity []categories.CategoryEntity
	for _, v := range category {
		categoryEntity = append(categoryEntity, CategoryToCategoryEntity(v))
	}
	return categoryEntity
}
