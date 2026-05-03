package handler

import "gorm.io/gorm"

type Handlers struct {
	Auth *AuthHandler
}

type AuthHandler struct {
	db *gorm.DB
}

func NewHandlers(db *gorm.DB) *Handlers  {
	return &Handlers {
		Auth: NewAuthHandler(db),
	}
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}
