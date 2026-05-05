package model

import "time"

type Post struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UserID     uint      `gorm:"not null"`
	CategoryID uint      `gorm:"not null"`
	Title      string    `gorm:"type:varchar(255);not null"`
	Content    string    `gorm:"type:text;not null"`
	Status     string    `gorm:"type:enum('draft','published');not null;default:'draft'"`
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time `gorm:"not null"`

	User User `gorm:"foreignKey:"UserID""`
	Category Category `gorm:"foreignKey":"CategoryID"`
}
