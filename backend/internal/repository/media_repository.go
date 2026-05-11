package repository

import (
	"gorm.io/gorm"
	"hobby-blog/internal/model"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{db: db}
}

func (r *MediaRepository) Create(media *model.MediaFile) error {
	return r.db.Create(media).Error
}
