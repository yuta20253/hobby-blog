package service

import (
	"gorm.io/gorm"
	"hobby-blog/internal/repository"
)
 "gorm.io/gorm"

type AuthService struct {
	repository  *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp() error {
	return s.repository.Create()
}
