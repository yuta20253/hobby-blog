package container

import (
	"gorm.io/gorm"
	"hobby-blog/internal/config"
	"hobby-blog/internal/handler"
	postApplication "hobby-blog/internal/post/application"
	userApplicationUsecase "hobby-blog/internal/user/application"
	"hobby-blog/internal/repository"
	"hobby-blog/internal/service"
	"hobby-blog/internal/storage"
	"hobby-blog/internal/uploader"
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
	authService := userApplicationUsecase.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	postRepo := repository.NewPostRepository(db)
	postService := postApplication.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	mypageService := service.NewMypageService(userRepo, postRepo)
	mypageHandler := handler.NewMypageHandler(mypageService)

	mediaRepo := repository.NewMediaRepository(db)

	st := storage.NewLocalStorage(cfg.UploadPath)
	upl := uploader.NewUploader(st)
	mediaService := service.NewMediaService(postRepo, mediaRepo, upl)
	mediaHandler := handler.NewMediaHandler(mediaService)

	return &Container{
		AuthHandler:   authHandler,
		PostHandler:   postHandler,
		MypageHandler: mypageHandler,
		MediaHandler:  mediaHandler,
	}
}
