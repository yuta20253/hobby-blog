package application

import postDomain "hobby-blog/internal/post/domain"

func ToPostDTO(post *postDomain.Post) *PostDTO {
	if post == nil {
		return nil
	}

	dto := NewPostDTO(*post)

	return &dto
}

func ToPostDTOs(posts []postDomain.Post) []PostDTO {
	dtos := make([]PostDTO, 0, len(posts))

	for _, post := range posts {
		dtos = append(dtos, NewPostDTO(post))
	}

	return dtos
}
