package model

import (
	"gorm.io/gorm"
	"hobby-blog/internal/domain/media"
)

type MediaFile struct {
	gorm.Model
	PostID   uint       `gorm:"not null" json:"post_id"`
	Type     media.Type `gorm:"type:varchar(20);not null" json:"type"`
	FilePath string     `gorm:"type:text;not null" json:"file_path"`
	FileName string     `gorm:"type:varchar(255);not null" json:"file_name"`
}
