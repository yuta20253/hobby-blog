package infrastructure

import (
	"context"
	"gorm.io/gorm"

	mediaDomain "hobby-blog/internal/media/domain"
)

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) mediaDomain.MediaRepository {
	return &mediaRepository{db: db}
}

func (r *mediaRepository) Create(ctx context.Context, media mediaDomain.Media) error {
	return r.db.WithContext(ctx).Create(&Media{
		PostID:   media.PostID,
		Type:     media.Type,
		FilePath: media.FilePath,
		FileName: media.FileName,
	}).Error
}
