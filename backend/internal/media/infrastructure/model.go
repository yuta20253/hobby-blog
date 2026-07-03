package infrastructure

import (
	"gorm.io/gorm"
	domainMedia "hobby-blog/internal/media/domain"
)

type Media struct {
	gorm.Model
	PostID   uint             `gorm:"not null" json:"post_id"`
	Type     domainMedia.Type `gorm:"type:varchar(20);not null" json:"type"`
	FilePath string           `gorm:"type:text;not null" json:"file_path"`
	FileName string           `gorm:"type:varchar(255);not null" json:"file_name"`
}
