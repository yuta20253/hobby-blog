package domain

import "context"

type MediaRepository interface {
	Create(ctx context.Context, media Media) error
}
