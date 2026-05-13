package service

import (
	"hobby-blog/internal/auth"
	"hobby-blog/internal/model"
	"hobby-blog/internal/pkg/password"
	"hobby-blog/internal/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

type AuthUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResult struct {
	User  AuthUserResponse `json:"user"`
	Token string           `json:"token"`
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(name, email, rawPassword string) (*AuthResult, error) {
	hashedPassword, err := password.Hash(rawPassword)

	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User: AuthUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Token: token,
	}, nil
}

func (s *AuthService) Login(email, rawPassword string) (*AuthResult, error) {
	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if err := password.Compare(user.PasswordHash, rawPassword); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User: AuthUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Token: token,
	}, nil
}

func (s *AuthService) GetUserByID(id uint) (*AuthUserResponse, error) {
	user, err := s.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return &AuthUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
