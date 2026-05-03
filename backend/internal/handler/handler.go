package handler

import "gorm.io/gorm"

type Handlers struct {
	Auth *AuthHandler
}

func NewHandlers(db *gorm.DB) *Handlers  {
	return &Handlers {
		Auth: NewAuthHandler(db),
	}
}
