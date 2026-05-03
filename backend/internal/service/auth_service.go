package service

import (
	"hobby-blog/internal/repository"
)

type AuthService struct {
	repo  *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp() error {
	return s.repo.Create()
}
