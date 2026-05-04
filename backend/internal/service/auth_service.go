package service

import (
	"hobby-blog/internal/repository"
	"hobby-blog/internal/model"
	"hobby-blog/internal/pkg/password"
)

type AuthService struct {
	repo  *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(name, email, rawPassword string) error {
	hashedPassword, err := password.Hash(rawPassword)

	if err != nil {
		return err
	}

	user := model.User{
		Name: name,
		Email: email,
		PasswordHash: hashedPassword,
	}

	return s.repo.Create(&user)
}
