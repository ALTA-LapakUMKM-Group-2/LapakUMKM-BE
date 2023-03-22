package delivery

import "lapakUmkm/features/categories"

type CategoryRequest struct {
	Category string `json:"category" form:"category"`
}

func CategoryRequestToCategoryEntity(categoryRequest *CategoryRequest) categories.CategoryEntity {
	return categories.CategoryEntity{
		Category: categoryRequest.Category,
	}
}
