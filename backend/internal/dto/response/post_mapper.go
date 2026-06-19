package response

import "hobby-blog/internal/model"

func NewPostResponse(post model.Post) PostResponse {
	return PostResponse{
		ID:    post.ID,
		Title: post.Title,
		User: PostUserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
		Category: CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
		MediaFiles: NewMediaFileResponses(post.MediaFiles),
	}
}

func NewPostResponses(posts []model.Post) []PostResponse {
	res := make([]PostResponse, 0, len(posts))

	for _, post := range posts {
		res = append(res, NewPostResponse(post))
	}

	return res
}

func NewPostDetailResponse(post model.Post) *PostDetailResponse {
	return &PostDetailResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		User: PostUserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
		Category: CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
		MediaFiles: NewMediaFileResponses(post.MediaFiles),
	}
}
