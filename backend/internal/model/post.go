package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID     uint      `gorm:"not null"`
	CategoryID uint      `gorm:"not null"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Content    string    `gorm:"type:text;not null"`
	Status     string    `gorm:"type:enum('draft','published');not null;default:'draft'"`

	User     User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
