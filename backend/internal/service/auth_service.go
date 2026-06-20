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
	serviceInput "hobby-blog/internal/service/input"
	"strings"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(input serviceInput.SignUpInput) (*response.AuthResult, error) {
	hashedPassword, err := password.Hash(input.Password)

	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:         input.Name,
		Email:        input.Email,
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

func (s *AuthService) Login(input serviceInput.LoginInput) (*response.AuthResult, error) {
	user, err := s.repo.FindByEmail(input.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrUnauthorized
		}
		return nil, err
	}

	if err := password.Compare(user.PasswordHash, input.Password); err != nil {
		return nil, appErrors.ErrUnauthorized
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
