package container

import (
	"gorm.io/gorm"
	"hobby-blog/internal/config"
	"hobby-blog/internal/handler"
	"hobby-blog/internal/repository"
	"hobby-blog/internal/service"
	"hobby-blog/internal/storage"
)

type Container struct {
	AuthHandler   *handler.AuthHandler
	PostHandler   *handler.PostHandler
	MypageHandler *handler.MypageHandler
	MediaHandler  *handler.MediaHandler
}

func NewContainer(db *gorm.DB) *Container {
	cfg := config.Load()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	mypageService := service.NewMypageService(userRepo, postRepo)
	mypageHandler := handler.NewMypageHandler(mypageService)

	mediaRepo := repository.NewMediaRepository(db)

	st := storage.NewLocalStorage(cfg.UploadPath)
	mediaService := service.NewMediaService(postRepo, mediaRepo, st)
	mediaHandler := handler.NewMediaHandler(mediaService)

	return &Container{
		AuthHandler:   authHandler,
		PostHandler:   postHandler,
		MypageHandler: mypageHandler,
		MediaHandler:  mediaHandler,
	}
}
