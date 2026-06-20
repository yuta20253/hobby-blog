package response

import "hobby-blog/internal/model"

func NewCategoryResponse(category model.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
