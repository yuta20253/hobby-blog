package presentation

import postApplication "hobby-blog/internal/post/application"

func NewPostResponse(post postApplication.PostDTO) PostResponse {
	return PostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		Category: CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
	}
}

func NewPostResponses(posts []postApplication.PostDTO) []PostResponse {
	responses := make([]PostResponse, 0, len(posts))

	for _, post := range posts {
		responses = append(responses, NewPostResponse(post))
	}

	return responses
}
