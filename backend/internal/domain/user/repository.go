package user

import "context"

type UserRepository interface {
	FindByID(ctx context.Context, id ID) (*User, error)
	FindByEmail(ctx context.Context, email Email) (*User, error)
	Create(ctx context.Context, user *User) error
}
