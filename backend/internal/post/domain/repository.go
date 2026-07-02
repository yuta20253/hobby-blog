package domain

import "context"

type PostRepository interface {
	Search(ctx context.Context, title, userName, category string, limit, offset int) ([]Post, error)
	GetByID(ctx context.Context, id uint) (*Post, error)
	Create(ctx context.Context, post Post) (*Post, error)
	Update(ctx context.Context, post Post) (*Post, error)
	Delete(ctx context.Context, id uint) error
	GetMyPostsByUserID(ctx context.Context, userID uint) ([]Post, error)
}
