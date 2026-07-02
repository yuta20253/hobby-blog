package infrastructure

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id"`
	CategoryID uint   `gorm:"not null" json:"category_id"`
	Title      string `gorm:"type:varchar(255);not null" json:"title"`
	Content    string `gorm:"type:text;not null" json:"content"`
	Status     string `gorm:"type:varchar(20);not null;default:'draft'" json:"status"`
}
