package response

import (
	postPresentationResponse "hobby-blog/internal/post/presentation"
	postInfrastructureModel "hobby-blog/internal/post/infrastructure"
)

func NewCategoryResponse(category postInfrastructureModel.Category) postPresentationResponse.CategoryResponse {
	return postPresentationResponse.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
