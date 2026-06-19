package service

import (
	"errors"
	"gorm.io/gorm"
	"hobby-blog/internal/dto/request"
	"hobby-blog/internal/dto/response"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/model"
	"hobby-blog/internal/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) SearchPosts(q request.SearchPostQuery) ([]response.PostResponse, error) {
	posts, err := s.repo.Search(q)

	if err != nil {
		return nil, err
	}

	res := make([]response.PostResponse, 0, len(posts))

	for _, post := range posts {
		res = append(res, response.PostResponse{
			ID:    post.ID,
			Title: post.Title,
			User: response.PostUserResponse{
				ID:    post.User.ID,
				Name:  post.User.Name,
				Email: post.User.Email,
			},
			Category: response.CategoryResponse{
				ID:   post.Category.ID,
				Name: post.Category.Name,
			},
			MediaFiles: response.NewMediaFileResponses(post.MediaFiles),
		})
	}

	return res, nil
}

func (s *PostService) GetPost(id uint) (*response.PostDetailResponse, error) {
	post, err := s.repo.Get(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		return nil, err
	}

	return &response.PostDetailResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		User: response.PostUserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
		Category: response.CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
		MediaFiles: response.NewMediaFileResponses(post.MediaFiles),
	}, nil
}

func (s *PostService) CreatePost(input request.CreatePostInput) error {
	return s.repo.Create(input)
}

func (s *PostService) UpdatePost(input request.UpdatePostInput) (*response.PostDetailResponse, error) {
	currentPost, err := s.repo.Get(input.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		return nil, err
	}

	if input.Status == model.StatusPublished && currentPost.Status != model.StatusDraft {
		return nil, appErrors.ErrInvalidInput
	}

	post, err := s.repo.Update(input)

	if err != nil {
		return nil, err
	}

	return &response.PostDetailResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		User: response.PostUserResponse{
			ID:   post.User.ID,
			Name: post.User.Name,
		},
		Category: response.CategoryResponse{
			ID:   post.Category.ID,
			Name: post.Category.Name,
		},
		MediaFiles: response.NewMediaFileResponses(post.MediaFiles),
	}, nil
}

func (s *PostService) DeletePost(id uint, userID uint) error {
	err := s.repo.Delete(id, userID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErrors.ErrNotFound
		}
		return err
	}

	return nil
}
