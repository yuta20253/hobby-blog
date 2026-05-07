package model

import (
	"gorm.io/gorm"
	"hobby-blog/internal/domain/post"
)

type Post struct {
	gorm.Model
	UserID     uint      `gorm:"not null"`
	CategoryID uint      `gorm:"not null"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Content    string    `gorm:"type:text;not null"`
	Status     post.Status    `gorm:"type:varchar(20);not null;default:'draft'"`

	User     User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
