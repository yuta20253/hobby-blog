package infrastructure

import (
	"context"

	"gorm.io/gorm"
	userDomain "hobby-blog/internal/user/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userDomain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*userDomain.User, error) {
	var m User
	if err := r.db.WithContext(ctx).First(&m, uint(id)).Error; err != nil {
		return nil, err
	}

	name, err := userDomain.NewName(m.Name)
	if err != nil {
		return nil, err
	}

	email, err := userDomain.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}

	return userDomain.Reconstruct(
		userDomain.ID(m.ID),
		name,
		email,
		m.PasswordHash,
	), nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email userDomain.Email) (*userDomain.User, error) {
	var m User
	if err := r.db.WithContext(ctx).Where("email = ?", email.String()).First(&m).Error; err != nil {
		return nil, err
	}

	mail, err := userDomain.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}

	name, err := userDomain.NewName(m.Name)
	if err != nil {
		return nil, err
	}

	return userDomain.Reconstruct(
		userDomain.ID(m.ID),
		name,
		mail,
		m.PasswordHash,
	), nil
}

func (r *userRepository) Create(ctx context.Context, u *userDomain.User) error {
	m := User{
		Name:         u.Name.String(),
		Email:        u.Email.String(),
		PasswordHash: u.PasswordHash,
	}

	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return nil
	}

	u.ID = userDomain.ID(m.ID)

	return nil
}
