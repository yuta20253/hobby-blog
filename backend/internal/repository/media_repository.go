package repository

import (
	"gorm.io/gorm"
	"hobby-blog/internal/model"
)

type mediaRepository struct {
	db *gorm.DB
}

type MediaRepository interface {
	Create(media *model.MediaFile) error
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{db: db}
}

func (r *mediaRepository) Create(media *model.MediaFile) error {
	return r.db.Create(media).Error
}
