package application

import (
    "context"
    "errors"

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

func (s *PostService) SearchPosts(ctx context.Context, query SearchPostQuery) ([]postDomain.Post, error) {
	posts, err := s.repo.Search(ctx, query.Title, query.UserName, query.Category, query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}
	responses := make([]postDomain.Post, 0, len(posts))
	for _, p := range posts {
		responses = append(responses, postDomain.Post{
			ID:      p.ID,
			Title:   p.Title,
			Content: p.Content,
		})
	}
	return responses, nil
}

func (s *PostService) GetPost(ctx context.Context, id uint) (*postDomain.Post, error) {
	post, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, appErrors.ErrNotFound
	}

	response := postDomain.Post{ID: post.ID, Title: post.Title, Content: post.Content}
	return &response, nil
}

func (s *PostService) CreatePost(ctx context.Context, input CreatePostInput) (*postDomain.Post, error) {
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

	response := postDomain.Post{ID: created.ID, Title: created.Title, Content: created.Content}
	return &response, nil
}

func (s *PostService) UpdatePost(ctx context.Context, input UpdatePostInput) (*postDomain.Post, error) {
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

	response := postDomain.Post{ID: result.ID, Title: result.Title, Content: result.Content}
	return &response, nil
}

func (s *PostService) DeletePost(ctx context.Context, id uint) error {
	post, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if post == nil {
		return appErrors.ErrNotFound
	}
	return s.repo.Delete(ctx, id)
}

func (s *PostService) PublishPost(ctx context.Context, id uint) error {
	post, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if post == nil {
		return appErrors.ErrNotFound
	}
	if post.Status != postDomain.StatusDraft {
		return errors.New("post is already published")
	}
	post.Status = postDomain.StatusPublished
	_, err = s.repo.Update(ctx, *post)
	return err
}
