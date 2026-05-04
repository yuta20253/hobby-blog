package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:255;unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
}
