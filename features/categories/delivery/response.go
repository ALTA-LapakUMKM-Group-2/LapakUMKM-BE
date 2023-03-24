package delivery

import "lapakUmkm/features/categories"

type CategoryResponse struct {
	Id       uint   `json:"id,omitempty"`
	Category string `json:"category" form:"category"`
}

func CategoryEntityToCategoryResponse(categoryEntity categories.CategoryEntity) CategoryResponse {
	return CategoryResponse{
		Id:       categoryEntity.Id,
		Category: categoryEntity.Category,
	}
}

func ListCategoryEntityToCategoryResponse(categoryEntity []categories.CategoryEntity) []CategoryResponse {
	var dataResponses []CategoryResponse
	for _, v := range categoryEntity {
		dataResponses = append(dataResponses, CategoryEntityToCategoryResponse(v))
	}
	return dataResponses
}
