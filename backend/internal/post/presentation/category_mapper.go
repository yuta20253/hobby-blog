package presentation

import postDomain "hobby-blog/internal/post/domain"

func NewCategoryResponse(category postDomain.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
