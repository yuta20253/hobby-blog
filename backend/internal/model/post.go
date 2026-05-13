package model

import (
	"gorm.io/gorm"
	"hobby-blog/internal/domain/post"
)

type Post struct {
	gorm.Model
	UserID     uint        `gorm:"not null" json:"user_id"`
	CategoryID uint        `gorm:"not null" json:"category_id"`
	Title      string      `gorm:"type:varchar(255);not null" json:"title"`
	Content    string      `gorm:"type:text;not null" json:"content"`
	Status     post.Status `gorm:"type:varchar(20);not null;default:'draft'" json:"status"`

	User       User        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Category   Category    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"category,omitempty"`
	MediaFiles []MediaFile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"media_files,omitempty"`
}
