package container

import (
	"gorm.io/gorm"
	"hobby-blog/internal/repository"
	"hobby-blog/internal/service"
	"hobby-blog/internal/handler"
)

type Container struct {
	AuthHandler *handler.AuthHandler
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	return &Container{
		AuthHandler: authHandler,
	}
}
