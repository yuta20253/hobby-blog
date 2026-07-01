package presentation

import postInfrastructureModel "hobby-blog/internal/post/infrastructure"

func NewPostResponse(post postInfrastructureModel.Post) PostResponse {
	return PostResponse{
		ID:         post.ID,
		Title:      post.Title,
		User:       NewPostUserResponse(post.User),
		Category:   NewCategoryResponse(post.Category),
		MediaFiles: NewMediaFileResponses(post.MediaFiles),
	}
}

func NewPostResponses(posts []postInfrastructureModel.Post) []PostResponse {
	res := make([]PostResponse, 0, len(posts))

	for _, post := range posts {
		res = append(res, NewPostResponse(post))
	}

	return res
}

func NewPostDetailResponse(post postInfrastructureModel.Post) *PostDetailResponse {
	return &PostDetailResponse{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		User:       NewPostUserResponse(post.User),
		Category:   NewCategoryResponse(post.Category),
		MediaFiles: NewMediaFileResponses(post.MediaFiles),
	}
}

func NewPostUserResponse(user postInfrastructureModel.User) PostUserResponse {
	return PostUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
