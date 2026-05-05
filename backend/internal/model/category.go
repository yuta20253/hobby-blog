package model

import "time"

type Category struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(50);unique;not null"`
}
