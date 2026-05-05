package service

import (
	"hobby-blog/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

type PostResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
}

type PostSearchQuery struct {
	Title    string
	UserName string
	Category string
	Limit    int
	Offset   int
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) SearchPosts(q PostSearchQuery) ([]PostResponse, error) {
	query := repository.PostSearchQuery{
		Title: q.Title,
		UserName: q.UserName,
		Category: q.Category,
		Limit: q.Limit,
		Offset: q.Offset,
	}

	posts, err := s.repo.Search(query)

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
