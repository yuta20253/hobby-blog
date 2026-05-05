package model

import "time"

type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(50);unique;not null"`
	CreatedAt time.Time `gorm:"not null"`
}
