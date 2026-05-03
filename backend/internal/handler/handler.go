package handler

import (
	"gorm.io/gorm"
	"hobby-blog/internal/repository"
)

type Handlers struct {
	Auth *AuthHandler
}

func NewHandlers(db *gorm.DB) *Handlers  {
	userRepo := repository.NewUserRepository(db)
	return &Handlers {
		Auth: NewAuthHandler(userRepo),
	}
}
