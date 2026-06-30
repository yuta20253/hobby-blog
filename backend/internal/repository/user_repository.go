package repository

import (
	"gorm.io/gorm"
	"hobby-blog/internal/model"
	"hobby-blog/internal/domain/user"
	"context"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	var m model.User
	if err := r.db.WithContext(ctx).First(&m, uint(id)).Error; err != nil {
		return nil, err
	}

	name, err := user.NewName(m.Name)
	if err != nil {
		return nil, err
	}

	email, err := user.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}

	return user.Reconstruct(
		user.ID(m.ID),
		name,
		email,
		m.PasswordHash,
	), nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	var m model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email.String()).First(&m).Error; err != nil {
		return nil, err
	}

	mail, err := user.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}

	name, err := user.NewName(m.Name)
	if err != nil {
		return nil, err
	}

	return user.Reconstruct(
		user.ID(m.ID),
		name,
		mail,
		m.PasswordHash,
	), nil
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	m := model.User{
		Name: u.Name.String(),
		Email: u.Email.String(),
		PasswordHash: u.PasswordHash,
	}

	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return nil
	}

	u.ID = user.ID(m.ID)

	return nil
}
