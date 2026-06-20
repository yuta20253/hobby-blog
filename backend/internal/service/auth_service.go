package service

import (
	"errors"
	"gorm.io/gorm"
	"hobby-blog/internal/auth"
	"hobby-blog/internal/dto/response"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/model"
	"hobby-blog/internal/pkg/password"
	"hobby-blog/internal/repository"
	"strings"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(name, email, rawPassword string) (*response.AuthResult, error) {
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
		if isDuplicateError(err) {
			return nil, appErrors.ErrConflict
		}
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &response.AuthResult{
		User:  response.NewAuthUserResponse(user),
		Token: token,
	}, nil
}

func (s *AuthService) Login(email, rawPassword string) (*response.AuthResult, error) {
	user, err := s.repo.FindByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrUnauthorized
		}
		return nil, err
	}

	if err := password.Compare(user.PasswordHash, rawPassword); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &response.AuthResult{
		User:  response.NewAuthUserResponse(user),
		Token: token,
	}, nil
}

func (s *AuthService) GetUserByID(id uint) (*response.AuthUserResponse, error) {
	user, err := s.repo.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		return nil, err
	}

	resp := response.NewAuthUserResponse(user)

	return &resp, nil
}

func isDuplicateError(err error) bool {
	// MySQL例（必要に応じて調整）
	return strings.Contains(err.Error(), "Duplicate entry")
}
