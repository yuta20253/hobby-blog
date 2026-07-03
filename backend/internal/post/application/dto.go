package application

import postDomain "hobby-blog/internal/post/domain"

type PostDTO struct {
	ID      uint
	Title   string
	Content string
	Status  string

	Category CategoryDTO
}

type UserDTO struct {
	ID    uint
	Name  string
	Email string
}

type CategoryDTO struct {
	ID   uint
	Name string
}

type MediaFileDTO struct {
	ID       uint
	Type     string
	FilePath string
	FileName string
}

func NewPostDTO(post postDomain.Post) PostDTO {
	return PostDTO{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		Status:  string(post.Status),
		Category: CategoryDTO{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
	}
}
