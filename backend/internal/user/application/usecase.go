package application

import (
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"

	"hobby-blog/internal/auth"
	appErrors "hobby-blog/internal/errors"
	"hobby-blog/internal/pkg/password"
	domainUser "hobby-blog/internal/user/domain"
)

type AuthService struct {
	repo domainUser.UserRepository
}

func NewAuthService(repo domainUser.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(ctx context.Context, input SignUpInput) (*AuthResult, error) {
	hashedPassword, err := password.Hash(input.Password)

	if err != nil {
		return nil, err
	}

	name, err := domainUser.NewName(input.Name)
	if err != nil {
		return nil, err
	}

	email, err := domainUser.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	user := domainUser.NewUser(
		name,
		email,
		hashedPassword,
	)

	if err := s.repo.Create(ctx, user); err != nil {
		if isDuplicateError(err) {
			return nil, appErrors.ErrConflict
		}
		return nil, err
	}

	token, err := auth.GenerateToken(uint(user.ID))
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User:  user,
		Token: token,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, input LoginInput) (*AuthResult, error) {
	email, err := domainUser.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrUnauthorized
		}
		return nil, err
	}

	if err := password.Compare(user.PasswordHash, input.Password); err != nil {
		return nil, appErrors.ErrUnauthorized
	}

	token, err := auth.GenerateToken(uint(user.ID))
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User:  user,
		Token: token,
	}, nil
}

func (s *AuthService) GetUserByID(ctx context.Context, id uint) (*domainUser.User, error) {
	user, err := s.repo.FindByID(ctx, domainUser.ID(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		return nil, err
	}

	return user, nil
}

func isDuplicateError(err error) bool {
	// MySQL例（必要に応じて調整）
	return strings.Contains(err.Error(), "Duplicate entry")
}
