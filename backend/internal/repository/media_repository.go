package repository

import (
	"gorm.io/gorm"
	postInfrastructureModel "hobby-blog/internal/post/infrastructure"
)

type mediaRepository struct {
	db *gorm.DB
}

type MediaRepository interface {
	Create(media *postInfrastructureModel.MediaFile) error
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db: db}
}

func (r *mediaRepository) Create(media *postInfrastructureModel.MediaFile) error {
	return r.db.Create(media).Error
}
