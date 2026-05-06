package service

import (
	"hobby-blog/internal/repository"
	"hobby-blog/internal/domain/post"
)

type PostService struct {
	repo *repository.PostRepository
}

type PostResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
}

type PostUserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PostDetailResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	User PostUserResponse `json:"user"`
	Category CategoryResponse `json:"category"`
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) SearchPosts(q post.SearchQuery) ([]PostResponse, error) {
	posts, err := s.repo.Search(q)

	if err != nil {
		return nil, err
	}

	res := make([]PostResponse, 0, len(posts))

	for _, p := range posts {
		res = append(res, PostResponse{
			ID: p.ID,
			Title: p.Title,
		})
	}

	return res, nil
}

func (s *PostService) GetPost(id uint) (*PostDetailResponse, error) {
	post, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return &PostDetailResponse{
		ID: post.ID,
		Title: post.Title,
		Content: post.Content,
		User: PostUserResponse{
			ID: post.User.ID,
			Name: post.User.Name,
		},
		Category: CategoryResponse{
			ID: post.Category.ID,
			Name: post.Category.Name,
		},
	}, nil
}

func (s *PostService) CreatePost(input post.CreateInput) error {
	return s.repo.Create(input)
}
