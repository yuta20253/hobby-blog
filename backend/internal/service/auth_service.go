package service

import (
	"hobby-blog/internal/repository"
	"hobby-blog/internal/model"
	"hobby-blog/internal/pkg/password"
	"hobby-blog/internal/auth"
)

type AuthService struct {
	repo  *repository.UserRepository
}

type SignUpResult struct {
	User model.User,
	Token string,
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(name, email, rawPassword string) (*SignUpResult, error) {
	hashedPassword, err := password.Hash(rawPassword)

	if err != nil {
		return err
	}

	user := model.User{
		Name: name,
		Email: email,
		PasswordHash: hashedPassword,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &SignUpResult{
		User: user,
		Token: token,
	}, nil
}
