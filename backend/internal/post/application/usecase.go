package application

import (
	"context"

	appErrors "hobby-blog/internal/errors"
	postDomain "hobby-blog/internal/post/domain"
	userDomain "hobby-blog/internal/user/domain"
)

type PostService struct {
	repo postDomain.PostRepository
}

func NewPostService(repo postDomain.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) SearchPosts(ctx context.Context, query SearchPostQuery) ([]PostDTO, error) {
	posts, err := s.repo.Search(ctx, query.Title, query.UserName, query.Category, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	return ToPostDTOs(posts), nil
}

func (s *PostService) GetPost(ctx context.Context, id uint) (*PostDTO, error) {
	post, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ToPostDTO(post), nil
}

func (s *PostService) CreatePost(ctx context.Context, input CreatePostInput) (*PostDTO, error) {
	post := postDomain.Post{
		Title:      input.Title,
		Content:    input.Content,
		CategoryID: input.CategoryID,
		UserID:     userDomain.ID(input.UserID),
		Status:     postDomain.StatusDraft,
	}
	created, err := s.repo.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return ToPostDTO(created), nil
}

func (s *PostService) UpdatePost(ctx context.Context, input UpdatePostInput) (*PostDTO, error) {
	post, err := s.repo.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, appErrors.ErrNotFound
	}
	updated := *post
	updated.Title = input.Title
	updated.Content = input.Content
	updated.CategoryID = input.CategoryID
	updated.Status = input.Status
	updated.UserID = userDomain.ID(input.UserID)
	result, err := s.repo.Update(ctx, updated)
	if err != nil {
		return nil, err
	}

	return ToPostDTO(result), nil
}

func (s *PostService) DeletePost(
	ctx context.Context,
	postID uint,
	userID uint,
) error {
	post, err := s.repo.GetByID(ctx, postID)
	if err != nil {
		return err
	}
	if post == nil {
		return appErrors.ErrNotFound
	}

	if uint(post.UserID) != userID {
		return appErrors.ErrForbidden
	}

	return s.repo.Delete(ctx, postID)
}
