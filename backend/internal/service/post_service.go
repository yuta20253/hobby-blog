package service

import (
	"errors"
	"hobby-blog/internal/domain/post"
	"hobby-blog/internal/dto"
	"hobby-blog/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

type PostDetailResponse struct {
	ID       uint                 `json:"id"`
	Title    string               `json:"title"`
	Content  string               `json:"content"`
	User     dto.PostUserResponse `json:"user"`
	Category dto.CategoryResponse `json:"category"`
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) SearchPosts(q post.SearchQuery) ([]dto.PostResponse, error) {
	posts, err := s.repo.Search(q)

	if err != nil {
		return nil, err
	}

	res := make([]dto.PostResponse, 0, len(posts))

	for _, p := range posts {
		res = append(res, dto.PostResponse{
			ID:    p.ID,
			Title: p.Title,
			User: dto.PostUserResponse{
				ID:    p.User.ID,
				Name:  p.User.Name,
				Email: p.User.Email,
			},
			Category: dto.CategoryResponse{
				ID:   p.Category.ID,
				Name: p.Category.Name,
			},
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
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		User: dto.PostUserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
		Category: dto.CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
	}, nil
}

func (s *PostService) CreatePost(input post.CreateInput) error {
	return s.repo.Create(input)
}

func (s *PostService) UpdatePost(input post.UpdateInput) (*PostDetailResponse, error) {
	currentPost, err := s.repo.Get(input.ID)

	if err != nil {
		return nil, err
	}

	if input.Status == post.StatusPublished && currentPost.Status != post.StatusDraft {
		return nil, errors.New("invalid status transition")
	}

	post, err := s.repo.Update(input)

	if err != nil {
		return nil, err
	}

	return &PostDetailResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		User: dto.PostUserResponse{
			ID:   post.User.ID,
			Name: post.User.Name,
		},
		Category: dto.CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
	}, nil
}

func (s *PostService) DeletePost(id uint, userID uint) error {
	return s.repo.Delete(id, userID)
}
